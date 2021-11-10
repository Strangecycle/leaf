package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"leaf/srv-user/conf"
	_ "leaf/srv-user/conf"
	"time"
)

func main() {
	reg := consul.NewRegistry(registry.Addrs(conf.GetConsulConf().Addr))

	service := micro.NewService(
		micro.Name("go.micro.leaf.user"),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)

	service.Init()

	// user.RegisterUserHandler(service.Server(), )

	logger.Fatal(service.Run().Error())
}
