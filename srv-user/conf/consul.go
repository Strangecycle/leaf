package conf

var consulConf Consul

type Consul struct {
	Addr string `yaml:"addr"`
}

func GetConsulConf() Consul {
	return consulConf
}
