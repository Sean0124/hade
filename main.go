package main

import (
	"coredemo/framework"
	"coredemo/framework/middleware"
	"log"
	"net/http"
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

	log.Fatal(server.ListenAndServe())
}
