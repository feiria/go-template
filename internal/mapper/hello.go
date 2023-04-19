package mapper

import (
	"go-template/global"
	"gorm.io/gorm"
)

type helloMapper struct {
	db *gorm.DB
}

func NewHelloMapper() HelloMapper {
	return &helloMapper{
		db: global.GetDB(),
	}
}

func (h helloMapper) GetHelloWorld() (string, error) {
	return "Hello World", nil
}
