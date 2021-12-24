package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"hade/framework/gin"
)

func FooControllerHandler(c *gin.Context) error {
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
		c.SetStatus(200).Json("ok")

		finish <- struct{}{}
	}()
	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.SetStatus(200).Json("panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.SetStatus(200).Json("time out")
		c.SetHasTimeout()
	}
	return nil
}

func UserLoginController(c *gin.Context) error {
	foo, _ := c.QueryString("foo", "def")
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)
	// 输出结果
	c.SetOkStatus().Json("ok, UserLoginController: " + foo)
	return nil
}

func SubjectListController(c *gin.Context) error {
	log.Println("SubjectListController")
	c.SetStatus(200).Json("ok, SubjectListController")
	return nil
}
func SubjectAddController(c *gin.Context) error {
	log.Println("SubjectAddController")
	c.SetStatus(200).Json("ok, SubjectAddController")
	return nil
}

func SubjectDelController(c *gin.Context) error {
	log.Println("SubjectDelController")
	c.SetStatus(200).Json("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *gin.Context) error {
	log.Println("SubjectUpdateController")
	c.SetStatus(200).Json("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *gin.Context) error {
	log.Println("SubjectGetController")
	c.SetStatus(200).Json("ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *gin.Context) error {
	log.Println("SubjectNameController")
	c.SetStatus(200).Json("ok, SubjectNameController")
	return nil
}
