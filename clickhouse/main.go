// main.go
package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


// SankeyNode ECharts Sankey 图需要的数据结构
type SankeyNode struct {
	Name string `json:"name"`
}

type SankeyLink struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Value  uint64 `json:"value"`
}

type SankeyData struct {
	Nodes []SankeyNode `json:"nodes"`
	Links []SankeyLink `json:"links"`
}

// 用于在 Go 中聚合 link 数据的辅助结构
type linkKey struct {
	Source string
	Target string
}

func main() {
	config, err := LoadConfig("./configs")
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	dsn := config.Database.ClickHouse.GetDSN()
	initDb(dsn)

	defer func() {
		db, err := DB.DB()
		if err != nil {
			db.Close()
		}
	}()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	// 1. 设置路由，提供前端页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	// 2. 设置 API 路由，提供桑基图数据
	r.GET("/api/sankey-data", getSankeyDataHandler)

	if err := r.Run(config.Server.Port); err != nil {
		log.Fatalf("无法启动服务器: %v", err)
	}
}

type Sankey struct {
	Path   string `json:"path" gorm:"column:source_path"`
	Totals uint64 `json:"totals" gorm:"column:totals"`
}

func RespError(ctx *gin.Context, error string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": "查询 ClickHouse 失败: " + error,
	})
	return
}

func getSankeyDataHandler(c *gin.Context) {
	var result = make([]Sankey, 0)

	tx := DB.WithContext(c).Raw(`SELECT source_path,countMerge(order_count) as totals FROM order_source_daily_metrics GROUP BY source_path;`).
		Scan(&result)

	if tx.Error != nil {
		RespError(c, tx.Error.Error())
		return
	}

	// --- 数据处理核心逻辑 ---
	// 聚合所有路径的 link 和 value
	linkValues := make(map[linkKey]uint64)
	// 使用 map 去重，得到所有 node
	nodeSet := make(map[string]struct{})

	for _, row := range result {
		// 将 'banner,detail,cart' 这种路径拆分为节点数组
		nodes := strings.Split(row.Path, ",")
		if len(nodes) < 1 {
			continue
		}

		// 将路径中的所有节点都加入到 nodeSet 中
		for _, node := range nodes {
			nodeSet[node] = struct{}{}
		}

		// 如果路径至少有两个节点 ，就好比 线必须有两头
		if len(nodes) < 2 {
			continue
		}

		// 从路径生成 link，例如 'a,b,c' -> (a->b), (b->c)
		for i := 0; i < len(nodes)-1; i++ {
			key := linkKey{Source: nodes[i], Target: nodes[i+1]}
			linkValues[key] += row.Totals // 累加相同 link 的订单数
		}
	}

	// 将 map 转换为 ECharts 需要的 slice 格式
	finalData := SankeyData{
		Nodes: make([]SankeyNode, 0, len(nodeSet)),
		Links: make([]SankeyLink, 0, len(linkValues)),
	}

	for nodeName := range nodeSet {
		finalData.Nodes = append(finalData.Nodes, SankeyNode{Name: nodeName})
	}

	for key, value := range linkValues {
		finalData.Links = append(finalData.Links, SankeyLink{Source: key.Source, Target: key.Target, Value: value})
	}
	// --- 数据处理结束 ---

	c.JSON(http.StatusOK, finalData)
}
