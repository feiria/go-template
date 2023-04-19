package main

import (
	"go-template/internal/handler"
	"go-template/internal/logic"
	"go-template/internal/mapper"
	"go-template/router"
)

func initServer() (*Server, func(), error) {
	helloMapper := mapper.NewHelloMapper()
	helloLogic := logic.NewHelloLogic(helloMapper)
	helloHandler := handler.NewHelloHandler(helloLogic)
	routerRouter := router.New(helloHandler)
	mainServer := New(routerRouter)
	return mainServer, func() {
	}, nil
}
