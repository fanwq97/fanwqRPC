package client

import (
	"context"
	"testing"
	"time"

	"github.com/fanwq97/fanwqRPC"
	"github.com/fanwq97/fanwqRPC/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {

	var ch = make(chan struct{})
	go func() {
		serverOpts := []fanwqRPC.ServerOption{
			fanwqRPC.WithAddress("127.0.0.1:8001"),
			fanwqRPC.WithNetwork("tcp"),
			fanwqRPC.WithSerializationType("msgpack"),
			fanwqRPC.WithTimeout(time.Millisecond * 2000),
		}
		s := fanwqRPC.NewServer(serverOpts...)
		if err := s.RegisterService("helloworld.Greeter", new(testdata.Service)); err != nil {
			panic(err)
		}

		go func() {
			s.Serve()
		}()

		<-ch
		s.Close()
	}()

	time.Sleep(1000 * time.Millisecond)

	opts := []Option{
		WithTarget("127.0.0.1:8001"),
		WithNetwork("tcp"),
		WithTimeout(2000 * time.Millisecond),
		WithSerializationType("msgpack"),
	}
	c := DefaultClient
	req := &testdata.HelloRequest{
		Msg: "hello",
	}
	rsp := &testdata.HelloReply{}

	err := c.Call(context.Background(), "/helloworld.Greeter/SayHello", req, rsp, opts...)

	close(ch)

	assert.Nil(t, err)
}
