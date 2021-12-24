package framework

// NewInstance 定义了如何创建一个新实例，所有服务容器的创建服务
type NewInstance func(...interface{}) (interface{}, error)


// ServiceProvider 定义一个服务提供者需要实现的接口
type ServiceProvider struct {
	Register(Container) NewInstance
	Boot(Container)  error
	isDefer() bool
	Params(Container) []interface{}
	Name() string
}