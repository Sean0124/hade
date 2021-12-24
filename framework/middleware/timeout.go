package middleware

import (
	"time"

	"github.com/Sean0124/hade/framework"
)

func TimeoutHandler(fun framework.ControllerHandler, d time.Duration) framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		time.Sleep(1 * time.Microsecond)
		return nil
	}
}
