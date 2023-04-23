package main

import (
	"time"

	"github.com/fanwq97/fanwqRPC"
	"github.com/fanwq97/fanwqRPC/plugin/consul"
	"github.com/fanwq97/fanwqRPC/testdata"
)

func main() {
	opts := []fanwqRPC.ServerOption{
		fanwqRPC.WithAddress("127.0.0.1:8000"),
		fanwqRPC.WithNetwork("tcp"),
		fanwqRPC.WithSerializationType("msgpack"),
		fanwqRPC.WithTimeout(time.Millisecond * 2000),
		fanwqRPC.WithSelectorSvrAddr("localhost:8500"),
		fanwqRPC.WithPlugin(consul.Name),
	}
	s := fanwqRPC.NewServer(opts...)
	if err := s.RegisterService("helloworld.Greeter", new(testdata.Service)); err != nil {
		panic(err)
	}
	s.Serve()
}