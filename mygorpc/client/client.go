package client

import (
	"context"
	"fmt"

	"github.com/lubanproj/gorpc/metadata"
	"github.com/lubanproj/gorpc/stream"

	"github.com/lubanproj/gorpc/codes"
	"github.com/lubanproj/gorpc/pool/connpool"
	"github.com/lubanproj/gorpc/protocol"
	"github.com/lubanproj/gorpc/selector"
	"github.com/lubanproj/gorpc/transport"
	"google.golang.org/protobuf/proto"

	"github.com/lubanproj/gorpc/codec"
)

type Client interface {
	Invoke(ctx context.Context, req, rsp interface{}, path string, opts ...ClientOption) error
}

type defaultClient struct {
	opts *ClientOptions
}

var DefaultClient = New()

var New = func() *defaultClient {
	return &defaultClient{
		opts: &ClientOptions{
			protocol: "proto",
		},
	}
}

// 使用gostruct定义结构体的方式，需要调用call
func (c *defaultClient) Call(ctx context.Context, servicePath string, req interface{}, rsp interface{},
	opts ...ClientOption) error {

	// reflection calls need to be serialized using msgpack
	callOpts := make([]ClientOption, 0, len(opts)+1)
	callOpts = append(callOpts, opts...)
	callOpts = append(callOpts, WithSerializationType(codec.MsgPack))

	// servicePath example : /helloworld.Greeter/SayHello
	err := c.Invoke(ctx, req, rsp, servicePath, callOpts...)
	if err != nil {
		return err
	}

	return nil
}

// 两种方式，不论是使用gostruct的反射方式还是proto代码生成，最终都会调用 invoke 函数。invoke 完成了一个客户端的完整动作
func (c *defaultClient) Invoke(ctx context.Context, req, rsp interface{}, path string, opts ...ClientOption) error {

}

func (c *defaultClient) invoke(ctx context.Context, req, rsp interface{}) error {

	//通过客户端透传下来的序列化类型参数，去获取 Serialization 对象，
	//然后通过 Serialization 对 request 进行序列化，会把请求体序列化成二进制数据
	serialization := codec.GetSerialization(c.opts.serializationType)
	payload, err := serialization.Marshal(req)
	if err != nil {
		return codes.NewFrameworkError(codes.ClientMsgErrorCode, "request marshal failed ...")
	}

	// assemble header
	// 将包头和包体拼一起
	request := addReqHeader(ctx, c, payload)

	reqbuf, err := proto.Marshal(request)
	if err != nil {
		return err
	}

	clientCodec := codec.GetCodec(c.opts.protocol)
	// 包头+包体序列化后拼上帧头，编码成二进制
	reqbody, err := clientCodec.Encode(reqbuf)
	if err != nil {
		return err
	}

	//底层 tcp 通信的能力是通过 transport 实现的，这里先创建一个client transport，

	clientTransport := c.NewClientTransport()
	clientTransportOpts := []transport.ClientTransportOption{
		transport.WithServiceName(c.opts.serviceName),
		transport.WithClientTarget(c.opts.target),
		transport.WithClientNetwork(c.opts.network),
		transport.WithClientPool(connpool.GetPool("default")),
		transport.WithSelector(selector.GetSelector(c.opts.selectorName)),
		transport.WithTimeout(c.opts.timeout),
	}
	//然后调用 transport 的 Send 函数往下游发送请求，会收到 server 返回的一个完整响应帧数据
	// 客户端将请求数据send到服务器，接受服务器返回的frame，这个frame包括帧头+包头+包体
	frame, err := clientTransport.Send(ctx, reqbody, clientTransportOpts...)
	if err != nil {
		return err
	}

	// 解码这里直接过滤了帧头，返回包头+包体
	rspbuf, err := clientCodec.Decode(frame)
	if err != nil {
		return err
	}

	// parse protocol header
	response := &protocol.Response{}
	if err = proto.Unmarshal(rspbuf, response); err != nil {
		return err
	}

	if response.RetCode != 0 {
		return codes.New(response.RetCode, response.RetMsg)
	}

	return serialization.Unmarshal(response.Payload, rsp)

}

func addReqHeader(ctx context.Context, client *defaultClient, payload []byte) *protocol.Request {
	clientStream := stream.GetClientStream(ctx)

	servicePath := fmt.Sprintf("/%s/%s", clientStream.ServiceName, clientStream.Method)
	md := metadata.ClientMetadata(ctx)

	// fill the authentication information
	for _, pra := range client.opts.perRPCAuth {
		authMd, _ := pra.GetMetadata(ctx)
		for k, v := range authMd {
			md[k] = []byte(v)
		}
	}

	request := &protocol.Request{
		ServicePath: servicePath,
		Payload:     payload,
		Metadata:    md,
	}

	return request
}

func (c *defaultClient) NewClientTransport() transport.ClientTransport {
	return transport.GetClientTransport(c.opts.protocol)
}
