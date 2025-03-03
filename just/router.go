package just

import (
	"log"
	"net/http"
	"strings"
)

// router 路由
type router struct {
	roots    map[string]*node       // 存储不同 HTTP 方法对应的 Trie 树根节点
	handlers map[string]HandlerFunc // 存储路由路径和对应的处理函数
}

// newRouter 创建 router 实例
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// parsePattern 分析路径，将路径字符串分割为路径片段
// 通配符处理: 若路径中包含通配符（*），则只保留通配符前的部分
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// addRoute 将路由添加到 Trie 树中，并保存对应的处理函数
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	log.Printf("Route %4s - %s", method, pattern)
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 not found:%s\n", c.Path)
	}
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	// 查找匹配的节点
	n := root.search(searchParts, 0)
	// 如果找到匹配节点
	if n != nil {
		// 解析路由参数
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			// 当前路径为通配符，并且后面有名字
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}
