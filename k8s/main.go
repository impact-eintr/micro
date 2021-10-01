package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("欢迎使用gin框架")
	s := []string{}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		s = append(s, "h")
		if len(s) > 1024 {
			log.Println("more than 1024")
		}
		c.String(http.StatusOK, s[0])
	})
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
	r.Run()
}
