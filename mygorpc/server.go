package mygorpc

import (
	"mygorpc/plugin"
	"os"
	"os/signal"
	"syscall"

	"github.com/lubanproj/gorpc/log"
	"github.com/lubanproj/gorpc/plugin/jaeger"
)

/*
要搭建 server 层，首先我们要明确 server 层需要支持哪些能力，其实 server 的核心就是提供服务请求的处理能力。
server 侧定义服务，发布服务，接收到服务的请求后，根据服务名和请求的方法名去路由到一个 handler 处理器，然后由 handler 处理请求，
得到响应，并且把响应数据发送给 client。
*/

type Server struct {
	opts     *ServerOptions     // 选项模型，用来透传业务自己指定的一些参数，比如服务监听的地址 address，网络类型 network 是 tcp 还是 udp，后端服务的超时时间 timeout 等。
	services map[string]Service //每个 Service 表示一个服务，一个 server 可以发布多个服务，用服务名 serviceName 作 map 的 key
	plugins  []plugin.Plugin    // Server 中添加 plugins 成员变量，它是一个插件数组。
	closing  bool               // whether the server is closing
}

func NewServer(opt ...ServerOption) *Server {
	s := &Server{
		opts:     &ServerOptions{},
		services: make(map[string]Service),
	}
	for _, o := range opt {
		o(s.opts)
	}
	//当调用 server.New 函数时，遍历插件 PluginMap，将所有插件 Plugin 添加到 plugins 中去
	for pluginName, plugin := range plugin.PluginMap {
		if !containPlugin(pluginName, s.opts.pluginNames) {
			continue
		}
		s.plugins = append(s.plugins, plugin)
	}
	return s
}

func containPlugin(pluginName string, plugins []string) bool {
	for _, plugin := range plugins {
		if pluginName == plugin {
			return true
		}
	}
	return false
}

func (s *Server) Serve() {
	//在调用 Server.Serve() 方法时，在 server 中的所有 service 提供服务之前，调用 InitPlugins 方法进行插件的配置初始化。
	err := s.InitPlugins()
	if err != nil {
		panic(err)
	}
	//遍历 service map 里面所有的 service，然后运行 service 的 Serve 方法
	for _, service := range s.services {
		go service.Serve(s.opts)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGSEGV)
	<-ch

	s.Close()
}

func (s *Server) InitPlugins() error {
	// init plugins
	for _, p := range s.plugins {

		switch val := p.(type) {

		case plugin.ResolverPlugin:
			var services []string
			services = append(services, s.service.Name())

			pluginOpts := []plugin.Option{
				plugin.WithSelectorSvrAddr(s.opts.selectorSvrAddr),
				plugin.WithSvrAddr(s.opts.address),
				plugin.WithServices(services),
			}
			if err := val.Init(pluginOpts...); err != nil {
				log.Errorf("resolver init error, %v", err)
				return err
			}

		case plugin.TracingPlugin:

			pluginOpts := []plugin.Option{
				plugin.WithTracingSvrAddr(s.opts.tracingSvrAddr),
			}

			tracer, err := val.Init(pluginOpts...)
			if err != nil {
				log.Errorf("tracing init error, %v", err)
				return err
			}

			s.opts.interceptors = append(s.opts.interceptors, jaeger.OpenTracingServerInterceptor(tracer, s.opts.tracingSpanName))

		default:

		}

	}

	return nil
}
