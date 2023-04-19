package main

import (
	"context"
	"flag"
	"fmt"
	"go-template/global"
	"go-template/pkg/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
	go gracefulShutdown()
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
