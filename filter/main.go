package main

import (
	"context"
	"fmt"

	"github.com/micro/examples/filter/network"
	proto "github.com/micro/examples/service/proto"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService()
	service.Init()

	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(
		// provide a context
		context.TODO(),
		// provide the request
		&proto.Request{Name: "John"},
		// set the filter
		network.Filter("micro"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
