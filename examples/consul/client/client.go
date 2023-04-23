package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fanwq97/fanwqRPC/client"
	"github.com/fanwq97/fanwqRPC/plugin/consul"
	"github.com/fanwq97/fanwqRPC/testdata"
)

func main() {
	opts := []client.Option{
		client.WithNetwork("tcp"),
		client.WithTimeout(2000 * time.Millisecond),
		client.WithSelectorName(consul.Name),
	}
	c := client.DefaultClient
	req := &testdata.HelloRequest{
		Msg: "hello",
	}
	rsp := &testdata.HelloReply{}

	consul.Init("localhost:8500")
	err := c.Call(context.Background(), "/helloworld.Greeter/SayHello", req, rsp, opts...)
	fmt.Println(rsp.Msg, err)
}
