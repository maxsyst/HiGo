package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{
			"name":  "GoLang",
			"value": "垒好",
		}
		//Output Json
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9000")
}
