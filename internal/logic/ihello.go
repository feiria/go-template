package logic

import "go-template/internal/types/response"

type HelloLogic interface {
	HelloAPI() (*response.HelloResponse, error)
}
