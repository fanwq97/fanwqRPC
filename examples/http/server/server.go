package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fanwq97/fanwqRPC"

	ghttp "github.com/fanwq97/fanwqRPC/http"
)

func init() {
	ghttp.HandleFunc("GET", "/hello", sayHello)
}

func main() {
	opts := []fanwqRPC.ServerOption{
		fanwqRPC.WithAddress("127.0.0.1:8000"),
		fanwqRPC.WithProtocol("http"),
		fanwqRPC.WithNetwork("tcp"),
		fanwqRPC.WithTimeout(time.Millisecond * 2000),
	}
	s := fanwqRPC.NewServer(opts...)
	s.ServeHttp()
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	w.Write([]byte("world"))
}
