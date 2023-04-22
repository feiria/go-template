#!/bin/sh

go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
go install github.com/swaggo/swag/cmd/swag
swag -v