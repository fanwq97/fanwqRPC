package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/fanwq97/fanwqRPC"
	"github.com/fanwq97/fanwqRPC/plugin/jaeger"
	"github.com/fanwq97/fanwqRPC/testdata"
)

func main() {

	pprof()

	opts := []fanwqRPC.ServerOption{
		fanwqRPC.WithAddress("127.0.0.1:8000"),
		fanwqRPC.WithNetwork("tcp"),
		fanwqRPC.WithSerializationType("msgpack"),
		fanwqRPC.WithTimeout(time.Millisecond * 2000),
		fanwqRPC.WithTracingSvrAddr("localhost:6831"),
		fanwqRPC.WithTracingSpanName("helloworld.Greeter"),
		fanwqRPC.WithPlugin(jaeger.Name),
	}
	s := fanwqRPC.NewServer(opts...)
	if err := s.RegisterService("helloworld.Greeter", new(testdata.Service)); err != nil {
		panic(err)
	}
	s.Serve()
}

func pprof() {
	go func() {
		http.ListenAndServe("0.0.0.0:8899", http.DefaultServeMux)
	}()
}