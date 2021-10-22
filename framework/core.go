package framework

import (
	"log"
	"net/http"
	"strings"
)

// framework core struct
type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler
}

// 初始化Core结构
func NewCore() *Core {
	// 初始化路由
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

// 匹配GET 方法, 增加路由规则
func (c *Core) Get(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// 匹配POST 方法, 增加路由规则
func (c *Core) Post(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// 匹配PUT 方法, 增加路由规则
func (c *Core) Put(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// 匹配DELETE 方法, 增加路由规则
func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers...); err != nil {
		log.Fatal("add router error: ", err)
	}

}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// framework core struct implementation Handler interface
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	// 封装自定义context
	ctx := NewContext(request, response)

	// 寻找路由
	node := c.FindRouteByRequest(request)
	if node == nil {
		// 如果没有找到，这里打印日志
		ctx.SetStatus(404).Json("not found")
		return
	}
	// 设置context中的handlers字段
	ctx.SetHandlers(node.handlers)

	params := node.parseParamsFromNode(request.URL.Path)
	ctx.SetParams(params)
	// 调用路由函数，如果返回err 代表存在内部错误，返回500状态码
	if err := ctx.Next(); err != nil {
		ctx.SetStatus(500).Json("inner error")
		return
	}
}

// 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) *node {
	// uri 和 method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		// 查找第二层map
		return methodHandlers.root.matchNode(uri)
	}
	return nil
}

// 注册中间件
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = middlewares
}
