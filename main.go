package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hade/framework"
	"hade/framework/middleware"
)

func main() {
	core := framework.NewCore()

	// core中使用use注册中间件
	core.Use(
		middleware.Test1(),
		middleware.Test2(),
	)

	// group中使用use注册中间件
	subjectApi := core.Group("/subject")
	subjectApi.Use(middleware.Recovery(), middleware.Test3())
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}
	go func() {
		server.ListenAndServe()
	}()
	// 当前的 Goroutine 等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前 Goroutine 等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
