package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Model
	LeafId       string    `json:"leaf_id"` // 个性化 Id
	Avatar       string    `json:"avatar"`
	Nickname     string    `json:"nickname"`
	Gender       int8      `gorm:"type:tinyint" json:"gender"`
	Phone        string    `gorm:"size:11" json:"phone"`
	Email        string    `json:"email"`
	Birth        string    `json:"birth"`
	Resident     string    `json:"resident"`                            // 常住地
	Hometown     string    `json:"hometown"`                            // 家乡
	RegisterTime time.Time `gorm:"autoCreateTime" json:"register_time"` // 注册时间
}

func (u User) TableName() string {
	return GetUserTableName()
}

func (u User) GetUserByPhone(phone string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("phone = ?", phone)
	}
}
