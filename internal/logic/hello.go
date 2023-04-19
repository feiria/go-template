package logic

import (
	"go-template/internal/mapper"
	"go-template/internal/types/response"
)

type helloLogic struct {
	helloMapper mapper.HelloMapper
}

func NewHelloLogic(helloMapper mapper.HelloMapper) HelloLogic {
	return &helloLogic{helloMapper: helloMapper}
}

func (l *helloLogic) HelloAPI() (resp *response.HelloResponse, err error) {
	if res, err := l.helloMapper.GetHelloWorld(); err != nil {
		return nil, err
	} else {
		return &response.HelloResponse{Word: res}, nil
	}
}
