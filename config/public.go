package config

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return db
}

func GetRDB() *redis.Client {
	return rdb
}

func GetAppPort() string {
	return appPort
}
