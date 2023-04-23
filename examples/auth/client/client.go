package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fanwq97/fanwqRPC/auth"
	"github.com/fanwq97/fanwqRPC/client"
	"github.com/fanwq97/fanwqRPC/testdata"
)

func main() {
	opts := []client.Option{
		client.WithTarget("127.0.0.1:8003"),
		client.WithNetwork("tcp"),
		client.WithTimeout(2000 * time.Millisecond),
		client.WithSerializationType("msgpack"),
		client.WithPerRPCAuth(auth.NewOAuth2ByToken("testToken")),
	}
	c := client.DefaultClient
	req := &testdata.HelloRequest{
		Msg: "hello",
	}
	rsp := &testdata.HelloReply{}
	err := c.Call(context.Background(), "/helloworld.Greeter/SayHello", req, rsp, opts...)
	fmt.Println(rsp.Msg, err)
}
