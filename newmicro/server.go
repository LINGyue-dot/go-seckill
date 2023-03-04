package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/v2"
	"github.com/micro/protoc-gen-micro/plugin/micro"
)

func main() {

	service := micro.NewService(
		// 服务注册唯一标识
		micro.Name("cap.imooc.server"),
	)

	service.Init()
	// 注册服务
	Register

}
