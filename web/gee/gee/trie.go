package gee

import "strings"

type node struct {
	// 待匹配路由，例如 /p/:lang 只有最后的子节点才有pattern
	pattern string
	// 路由中的一部分，例如 :lang
	part string
	// 子节点，例如 [ doc , tutorial , intro]
	children []*node
	// 非精确匹配, part 含有 : 或 * 时为 true
	isWild bool
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {

	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)

	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

// insert 递归查找每一层的节点，如果没有匹配到当前 part 的节点，则新建一个
// `/p/:lang/doc` 只有在第三层节点，即 doc 节点，pattern 才会设置为 /p/:lang/doc 。 p 和 :lang 节点的 pattern 属性皆为空。
// height 表示为 路径层级
func (n *node) insert(pattern string, parts []string, height int) {
	// 最后一层
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]

	child := n.matchChild(part)
	if child == nil { // 没有节点命中，说明是新的节点
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*', // 非精确匹配
		}

		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

// search 递归搜索路径
// parts 表示所要搜索的路径
// 递归查询每一层的节点，匹配到了*、匹配失败、匹配到了第len(parts)层节点就退出
func (n *node) search(parts []string, height int) *node {
	// 最后一层 或者 当前节点是 * (规定 * 只有一个，并且只能是最后一个)
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// 当匹配结束时，我们可以使用 n.pattern == "" 来判断路由规则是否匹配成功。
		// 例如，/p/python虽能成功匹配到:lang，但 :lang的 pattern 值为空，因此匹配失败。
		if n.pattern == "" { // 最后一个节点的 pattern 值不为空,说明匹配的花，必须匹配到最后节点才行，不然就是匹配失败
			return nil
		}

		return n
	}

	part := parts[height]

	// 当前节点下的所有子节点
	children := n.matchChildren(part)

	for _, child := range children {
		// 当前子节点下的每一个节点 都搜索子节点
		result := child.search(parts, height+1)

		if result != nil {
			return result
		}
	}

	return nil
}
