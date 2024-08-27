package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 引擎
	router := gin.Default()

	// 定义一个 GET 请求的路由，路径为 /hello
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// 启动服务器，监听在 8080 端口
	router.Run(":9001")
}
