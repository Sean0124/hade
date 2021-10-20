package main

import (
	"context"
	"coredemo/framework"
	"fmt"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(2*time.Second))
	defer cancel()

	// mu := sync.Mutex{}
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// Do real action
		time.Sleep(1 * time.Second)
		c.Json(200, "ok")

		finish <- struct{}{}
	}()
	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout()
	}
	return nil
}

func UserLoginController(c *framework.Context) error {
	log.Println("UserLoginController")
	c.Json(200, "ok, UserLoginController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	log.Println("SubjectListController")
	c.Json(200, "ok, SubjectListController")
	return nil
}
func SubjectAddController(c *framework.Context) error {
	log.Println("SubjectAddController")
	c.Json(200, "ok, SubjectAddController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	log.Println("SubjectDelController")
	c.Json(200, "ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	log.Println("SubjectUpdateController")
	c.Json(200, "ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	log.Println("SubjectGetController")
	c.Json(200, "ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	log.Println("SubjectNameController")
	c.Json(200, "ok, SubjectNameController")
	return nil
}
