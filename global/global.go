package global

import (
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	ConcurrencyControl = &singleflight.Group{}
	Log                *zap.Logger
)

func GetDB() *gorm.DB {
	return DB
}

func GetLog() *zap.Logger {
	return Log
}
