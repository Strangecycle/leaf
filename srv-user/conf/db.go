package conf

var dbConf DB

type DB struct {
	Host         string
	Port         string
	Config       string
	DbName       string
	Username     string
	Password     string
	MaxIdleConns int
	MaxOpenConns int
	DebugMode    bool
	LogMode      bool
}

func GetDBConf() DB {
	return dbConf
}
