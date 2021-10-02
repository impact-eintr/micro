package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("欢迎使用gin框架")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})
	started := time.Now()
	r.GET("/healthz", func(c *gin.Context) {
		duration := time.Now().Sub(started)
		if duration.Seconds() < 20 {
			c.JSON(http.StatusServiceUnavailable, nil)
		} else {
			c.JSON(http.StatusOK, nil)
		}
	})
	r.Run()
}
