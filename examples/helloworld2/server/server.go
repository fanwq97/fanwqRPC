package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fanwq97/fanwqRPC"
	"github.com/fanwq97/fanwqRPC/examples/helloworld2/helloworld"
)

type greeterService struct{}

func (g *greeterService) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Println("recv Msg : ", req.Msg)
	rsp := &helloworld.HelloReply{
		Msg: req.Msg + " world",
	}
	return rsp, nil
}

func main() {
	opts := []fanwqRPC.ServerOption{
		fanwqRPC.WithAddress("127.0.0.1:8000"),
		fanwqRPC.WithNetwork("tcp"),
		fanwqRPC.WithProtocol("proto"),
		fanwqRPC.WithTimeout(time.Millisecond * 2000),
	}
	s := fanwqRPC.NewServer(opts...)
	helloworld.RegisterService(s, &greeterService{})
	s.Serve()
}
