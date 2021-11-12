package common

import (
	"fmt"
	"github.com/micro/go-micro/v2/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"leaf/srv-user/conf"
	"leaf/srv-user/models"
)

var gormDB *gorm.DB

func init() {
	c := conf.GetDBConf()
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?%v",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DbName,
		c.Config,
	)

	mysqlConf := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         188,   // string 类型默认大小
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConf))
	if err != nil {
		logger.Fatal("failed to open database: ", err.Error())
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)

	if c.DebugMode {
		db = db.Debug()
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		logger.Fatal("failed to auto migrate: ", err.Error())
	}

	gormDB = db
}

func GetDB() *gorm.DB {
	return gormDB
}
