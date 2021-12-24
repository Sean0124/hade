package framework

import "sync"

// Container是一个服务容器，提供绑定服务和获取服务的功能
type Container interface {
	// Bind绑定一个服务提供者，如果关键字凭证存在，会进行替换操作，返回error
	Bind(provider ServiceProvider) error
	// IsBind关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool
	// Make根据关键字凭证获取一个服务
	Make(key string) (interface{}, error)
	// MustMake根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会panic
	// 所以在使用这个接口的时候请保证服务已经为这个关键字凭证绑定了服务提供者
	MustMake(key string) interface{}
	// MakeNew根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务器提供者注册的启动函数和传递的params参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// HadeContainer是服务容器的具体实现
type HadeContainer struct {
	Container //强烈要求HadeContainer实现Container接口
	// providers 存储注册的服务提供者，key为字符凭证
	providers map[string]ServiceProvider
	// instances 存储具体的实例，key为字符凭证
	instances map[string]interface{}
	// lock用于锁住对容器的变更操作
	lock sync.RWMutex
}

func (hade *HadeContainer) Bind(provider ServiceProvider) error {
	hade.lock.Lock()
	defer hade.lock.Unlock()
	key := provider.Name()

	hade.providers[key] = provider

	if provider.IsDefer() = false {
		
	}

}
