package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"strings"
)

// RouteInfo 用于存储单个路由的信息
type RouteInfo struct {
	Method      string
	Path        string
	Middlewares []string
	Handler     string
}

// GroupInfo 用于存储路由分组信息
type GroupInfo struct {
	Name        string     // 变量名，如 v1
	Prefix      string     // 完整的路径前缀，如 /api/v1
	Middlewares []string   // 该分组应用的中间件
	Parent      *GroupInfo // 父分组
}

// Analyser 是我们的核心分析器结构
type Analyser struct {
	fset              *token.FileSet
	routes            []*RouteInfo
	groups            map[string]*GroupInfo // key是group变量名，如 "v1"
	globalMiddlewares []string
	engineVar         string // Gin引擎的变量名，通常是 "r"
}

func NewAnalyser() *Analyser {
	return &Analyser{
		fset:   token.NewFileSet(),
		routes: make([]*RouteInfo, 0),
		groups: make(map[string]*GroupInfo),
	}
}

// ParseFile 解析单个 Go 文件
func (a *Analyser) ParseFile(filePath string) error {
	node, err := parser.ParseFile(a.fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	ast.Inspect(node, a.nodeVisitor)
	return nil
}

// nodeVisitor 是 AST 的核心访问者函数
func (a *Analyser) nodeVisitor(n ast.Node) bool {
	// 查找并处理赋值语句，特别是 a. gin引擎的赋值 b. 路由分组的赋值
	if assign, ok := n.(*ast.AssignStmt); ok {
		a.handleAssignment(assign)
		return true // 继续遍历子节点
	}

	// 查找方法调用，如 r.GET, r.Use
	if call, ok := n.(*ast.CallExpr); ok {
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
			receiverName := getReceiverName(sel)
			methodName := sel.Sel.Name

			// 1. 处理中间件 r.Use(...) 或 v1.Use(...)
			if methodName == "Use" {
				a.handleUseMiddleware(receiverName, call)
			}

			// 2. 处理HTTP方法路由 r.GET(...) 或 v1.GET(...)
			if a.isHTTPMethod(methodName) {
				a.handleRoute(receiverName, methodName, call)
			}
		}
	}
	return true
}

// handleAssignment 专门处理赋值语句
func (a *Analyser) handleAssignment(assign *ast.AssignStmt) {
	if len(assign.Lhs) != 1 || len(assign.Rhs) != 1 {
		return
	}
	// 获取左侧变量名
	varName, ok := assign.Lhs[0].(*ast.Ident)
	if !ok {
		return
	}
	// 获取右侧的函数调用
	call, ok := assign.Rhs[0].(*ast.CallExpr)
	if !ok {
		return
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	// Case A: 寻找 gin 引擎, r := gin.Default()
	if x, ok := sel.X.(*ast.Ident); ok && x.Name == "gin" && (sel.Sel.Name == "New" || sel.Sel.Name == "Default") {
		a.engineVar = varName.Name
		log.Printf("发现 Gin 引擎变量: %s\n", a.engineVar)
		if sel.Sel.Name == "Default" {
			a.globalMiddlewares = append(a.globalMiddlewares, "gin.Logger", "gin.Recovery")
		}
		// 把引擎本身也当作一个没有前缀的“根分组”
		a.groups[a.engineVar] = &GroupInfo{Name: a.engineVar, Prefix: "/"}
		return
	}

	// Case B: 寻找路由分组, v1 := r.Group("/api/v1")
	if sel.Sel.Name == "Group" {
		receiverName := getReceiverName(sel)
		if _, isKnownGroup := a.groups[receiverName]; isKnownGroup {
			prefix := a.extractStringLit(call.Args[0])
			parentGroup := a.groups[receiverName]

			newGroup := &GroupInfo{
				Name:        varName.Name,
				Prefix:      path.Join(parentGroup.Prefix, prefix),
				Middlewares: []string{},
				Parent:      parentGroup,
			}
			a.groups[newGroup.Name] = newGroup
			log.Printf("发现路由分组: %s = %s.Group(\"%s\")\n", newGroup.Name, receiverName, prefix)
		}
	}
}

// handleUseMiddleware 处理 r.Use() or v1.Use()
func (a *Analyser) handleUseMiddleware(receiverName string, call *ast.CallExpr) {
	middlewares := a.extractHandlers(call.Args)
	if receiverName == a.engineVar {
		// 全局中间件
		a.globalMiddlewares = append(a.globalMiddlewares, middlewares...)
		log.Printf("发现全局中间件: %v\n", middlewares)
	} else if group, ok := a.groups[receiverName]; ok {
		// 分组中间件
		group.Middlewares = append(group.Middlewares, middlewares...)
		log.Printf("发现分组 [%s] 中间件: %v\n", receiverName, middlewares)
	}
}

// handleRoute 处理具体的路由定义
func (a *Analyser) handleRoute(receiverName, methodName string, call *ast.CallExpr) {
	group, ok := a.groups[receiverName]
	if !ok {
		return // 不是已知的 group 或 engine，忽略
	}
	if len(call.Args) < 1 {
		return
	}
	routePath := a.extractStringLit(call.Args[0])
	handlers := a.extractHandlers(call.Args[1:])

	// 分离中间件和最终的处理器
	inlineMiddlewares := make([]string, 0)
	handlerName := "N/A"
	if len(handlers) > 0 {
		inlineMiddlewares = handlers[:len(handlers)-1]
		handlerName = handlers[len(handlers)-1]
	}

	// --- 收集所有中间件 ---
	allMiddlewares := make([]string, 0)
	// 1. 全局中间件
	allMiddlewares = append(allMiddlewares, a.globalMiddlewares...)

	// 2. 从当前分组回溯到根，收集所有父分组的中间件
	groupChain := make([]*GroupInfo, 0)
	for g := group; g != nil && g.Name != a.engineVar; g = g.Parent {
		groupChain = append(groupChain, g)
	}
	// 倒序添加，以保证父中间件在前
	for i := len(groupChain) - 1; i >= 0; i-- {
		allMiddlewares = append(allMiddlewares, groupChain[i].Middlewares...)
	}

	// 3. 路由行内中间件
	allMiddlewares = append(allMiddlewares, inlineMiddlewares...)

	route := &RouteInfo{
		Method:      strings.ToUpper(methodName),
		Path:        path.Join(group.Prefix, routePath),
		Middlewares: allMiddlewares,
		Handler:     handlerName,
	}
	a.routes = append(a.routes, route)
	log.Printf("发现路由: %s %s, 中间件: %v, 处理器: %s\n", route.Method, route.Path, route.Middlewares, route.Handler)
}

// GenerateMarkdown 生成 Markdown 文件
func (a *Analyser) GenerateMarkdown(outputFile string) error {
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	var sb strings.Builder
	sb.WriteString("# API 中间件文档\n\n")
	sb.WriteString("| 方法 | 路径 | 中间件执行链 |\n")
	sb.WriteString("|:---|:---|:---|\n")

	for _, route := range a.routes {
		middlewareChain := "无"
		if len(route.Middlewares) > 0 {
			cleanedMws := make([]string, len(route.Middlewares))
			for i, mw := range route.Middlewares {
				// 清理函数名，去掉可能存在的()
				cleanedMws[i] = strings.TrimSuffix(mw, "()")
			}
			middlewareChain = strings.Join(cleanedMws, " -> ")
		}

		// --- 这里是修复的地方 ---
		// 删除了多余的 route.Middlewares 参数
		line := fmt.Sprintf("| %s | `%s` | `%s` |\n", route.Method, route.Path, middlewareChain)
		sb.WriteString(line)
	}

	_, err = f.WriteString(sb.String())
	return err
}

// --- AST 辅助函数 (与之前版本相同) ---
func getReceiverName(sel *ast.SelectorExpr) string {
	if ident, ok := sel.X.(*ast.Ident); ok {
		return ident.Name
	}
	return ""
}
func (a *Analyser) extractHandlers(args []ast.Expr) []string {
	var handlers []string
	for _, arg := range args {
		switch expr := arg.(type) {
		case *ast.Ident:
			handlers = append(handlers, expr.Name)
		case *ast.CallExpr:
			// 尝试从函数调用中提取函数名
			var funcName string
			// 可能是 simpleIdentifier()
			if ident, ok := expr.Fun.(*ast.Ident); ok {
				funcName = ident.Name
			}
			// 可能是 package.Selector()
			if sel, ok := expr.Fun.(*ast.SelectorExpr); ok {
				if x, ok := sel.X.(*ast.Ident); ok {
					funcName = x.Name + "." + sel.Sel.Name
				}
			}
			if funcName != "" {
				handlers = append(handlers, funcName+"()")
			}
		}
	}
	return handlers
}
func (a *Analyser) extractStringLit(expr ast.Expr) string {
	if lit, ok := expr.(*ast.BasicLit); ok && lit.Kind == token.STRING {
		return strings.Trim(lit.Value, `"`)
	}
	return ""
}
func (a *Analyser) isHTTPMethod(name string) bool {
	methods := map[string]bool{"GET": true, "POST": true, "PUT": true, "DELETE": true, "PATCH": true, "OPTIONS": true, "HEAD": true, "Any": true}
	return methods[name]
}

func main() {
	analyser := NewAnalyser()

	// 单次解析即可
	if err := analyser.ParseFile("main.go"); err != nil {
		log.Fatalf("解析文件失败: %v", err)
	}

	if err := analyser.GenerateMarkdown("docs/api_docs.md"); err != nil {
		log.Fatalf("生成Markdown失败: %v", err)
	}

	fmt.Println("API文档 'api_docs.md' 已成功生成!")
}
