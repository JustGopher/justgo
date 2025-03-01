package just

import (
	"net/http"
)

type HandlerFunc func(*Context)

// Engine 实现 ServeHTTP 接口
type Engine struct {
	router *router
}

// New Engine的构造器
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

// addRoute 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET 添加Get请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodGet, pattern, handler)
}

// POST 添加Post请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute(http.MethodPost, pattern, handler)
}

// Run 启动http服务
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}
