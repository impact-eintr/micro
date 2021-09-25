package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

var count int32 = 0

func main() {
	fmt.Println("欢迎使用gin框架")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		atomic.AddInt32(&count, 1)
		log.Println(count)
		c.String(http.StatusOK, "Hello, 世界")
	})
	r.Run()
}
