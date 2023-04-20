package main

import (
	"context"
	"flag"
	"fmt"
	_ "go-template/docs"
	"go-template/global"
	"go-template/internal/logic"
	"go-template/pkg/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//	@title			接口文档
//	@version		1.0.0
//	@description	接口文档
//	@host			localhost:8866
//	@BasePath		/api
//	@termsOfService	https://github.com/feiria/go-template

var (
	env        string
	signalChan = make(chan os.Signal, 1)
	server     *http.Server
)

func init() {
	flag.StringVar(&env, "env", "dev", "obtain the current environment parameters e.g. 'prod','test','dev'")
}

func main() {
	flag.Parse()
	config.Init(env)

	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// 监听中止信号
	go gracefulShutdown()
	// 开启定时任务
	go logic.CronExample()

	app, cancel, err := initServer()
	defer cancel()
	if err != nil {
		panic(err)
	}

	server = app.Start()
	global.Log.Info(fmt.Sprintf("API Server start at %s", server.Addr))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		global.Log.Fatal(fmt.Sprintf("start gamefi openapi server failed, err: %s", err))
	}
}

func gracefulShutdown() {
	select {
	case msg := <-signalChan:
		switch msg {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM:
			global.Log.Info("API Server Shutdown......")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				global.Log.Error(fmt.Sprintf("API Server Shutdown Error: %+v", err))
			}
		}
		os.Exit(0)
	}
}
