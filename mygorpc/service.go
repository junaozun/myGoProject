package mygorpc

import "context"

//Service 的接口定义了每个服务需要提供的通用能力，包括 Register （处理函数 Handler 的注册）、提供服务 Serve，服务关闭 Close 等方法
type Service interface {
	Register(string, Handler)
	Serve(*ServerOptions)
	Close()
	Name() string
}

//是 Service 接口的具体实现。它的核心是 handlers 这个 map，每一类请求会分配一个 Handler 进行处理
type service struct {
	svr         interface{}        // server
	ctx         context.Context    // 每一个 service 一个上下文进行管理
	cancel      context.CancelFunc // context 的控制器
	serviceName string             // 服务名
	handlers    map[string]Handler // 方法名：Handler
	opts        *ServerOptions     //参数选项

	closing bool // whether the service is closing
}

//构建 transport ，监听客户端请求，根据请求的服务名 serviceName 和请求的方法名 methodName
//调用相应的 handler 去处理请求，然后进行回包。
func (s *service) Serve(options *ServerOptions) {

}
