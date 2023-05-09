//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-template/internal/handler"
	"go-template/internal/logic"
	"go-template/internal/mapper"
	"go-template/router"
)

func initWire() (s *Server, cancel func(), err error) {
	panic(wire.Build(
		mapper.Set,
		logic.Set,
		handler.Set,
		New,
		router.New,
	))
	return
}
