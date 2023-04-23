package main

import (
	"context"
	"errors"
	"time"

	"github.com/fanwq97/fanwqRPC"
	"github.com/fanwq97/fanwqRPC/auth"
	"github.com/fanwq97/fanwqRPC/log"
	"github.com/fanwq97/fanwqRPC/metadata"
	"github.com/fanwq97/fanwqRPC/testdata"
)

func main() {

	af := func(ctx context.Context) (context.Context, error) {
		md := metadata.ServerMetadata(ctx)

		if len(md) == 0 {
			return ctx, errors.New("token nil")
		}
		v := md["authorization"]
		log.Debug("token : ", string(v))
		if string(v) != "Bearer testToken" {
			return ctx, errors.New("token invalid")
		}
		return ctx, nil
	}

	opts := []fanwqRPC.ServerOption{
		fanwqRPC.WithAddress("127.0.0.1:8003"),
		fanwqRPC.WithNetwork("tcp"),
		fanwqRPC.WithSerializationType("msgpack"),
		fanwqRPC.WithTimeout(time.Millisecond * 2000),
		fanwqRPC.WithInterceptor(auth.BuildAuthInterceptor(af)),
	}
	s := fanwqRPC.NewServer(opts...)
	if err := s.RegisterService("/helloworld.Greeter", new(testdata.Service)); err != nil {
		panic(err)
	}
	s.Serve()
}
