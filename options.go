package fanwqRPC

import (
	"time"

	"github.com/fanwq97/fanwqRPC/interceptor"
)

// ServerOptions defines the server serve parameters
type ServerOptions struct {
	address           string        // 监听地址, e.g. :( ip://127.0.0.1:8080、 dns://www.google.com)
	network           string        // 网络协议类型, e.g. : tcp、udp
	protocol          string        // 序列化的数据 e.g. : proto、json
	timeout           time.Duration // 超时时间
	serializationType string        // 序列化类型, default: proto

	selectorSvrAddr string                          // 使用第三方服务发现插件时所需的服务发现服务器地址
	tracingSvrAddr  string                          // 使用第三方跟踪插件时所需的跟踪插件服务器地址
	tracingSpanName string                          // 使用第三方跟踪插件时所需的跟踪跨度名称
	pluginNames     []string                        // 插件名称
	interceptors    []interceptor.ServerInterceptor // 拦截器
}

type ServerOption func(*ServerOptions)

func WithAddress(address string) ServerOption {
	return func(o *ServerOptions) {
		o.address = address
	}
}

func WithNetwork(network string) ServerOption {
	return func(o *ServerOptions) {
		o.network = network
	}
}

func WithProtocol(protocol string) ServerOption {
	return func(o *ServerOptions) {
		o.protocol = protocol
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(o *ServerOptions) {
		o.timeout = timeout
	}
}

func WithSerializationType(serializationType string) ServerOption {
	return func(o *ServerOptions) {
		o.serializationType = serializationType
	}
}

func WithSelectorSvrAddr(addr string) ServerOption {
	return func(o *ServerOptions) {
		o.selectorSvrAddr = addr
	}
}

func WithPlugin(pluginName ...string) ServerOption {
	return func(o *ServerOptions) {
		o.pluginNames = append(o.pluginNames, pluginName...)
	}
}

func WithInterceptor(interceptors ...interceptor.ServerInterceptor) ServerOption {
	return func(o *ServerOptions) {
		o.interceptors = append(o.interceptors, interceptors...)
	}
}

func WithTracingSvrAddr(addr string) ServerOption {
	return func(o *ServerOptions) {
		o.tracingSvrAddr = addr
	}
}

func WithTracingSpanName(name string) ServerOption {
	return func(o *ServerOptions) {
		o.tracingSpanName = name
	}
}
