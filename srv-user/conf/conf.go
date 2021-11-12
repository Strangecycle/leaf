package conf

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/encoder/yaml"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/logger"
)

const configFilePath = "./conf.yaml"

func init() {
	// GoLand 启动服务时必须将 Working Directory 设置成服务根目录
	// 否则会报找不到根路径
	encoder := yaml.NewEncoder()
	c, err := config.NewConfig(
		config.WithSource(
			file.NewSource(
				file.WithPath(configFilePath),
				source.WithEncoder(encoder)),
		),
	)

	if err != nil {
		logger.Fatal("failed to load config: ", err.Error())
	}

	if err := c.Load(file.NewSource(file.WithPath(configFilePath))); err != nil {
		logger.Fatal("failed to load config: ", err.Error())
	}

	err = c.Get("consul").Scan(&consulConf)
	err = c.Get("database").Scan(&dbConf)

	if err != nil {
		logger.Fatal("failed to load config: ", err.Error())
	}
}
