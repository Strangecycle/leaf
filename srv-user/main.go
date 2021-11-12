package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	_ "leaf/srv-user/common"
	"leaf/srv-user/conf"
	_ "leaf/srv-user/conf"
	"leaf/srv-user/handler"
	"leaf/srv-user/proto/out/user"
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

	_ = user.RegisterUserHandler(service.Server(), new(handler.User))

	logger.Fatal(service.Run().Error())
}
