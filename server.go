package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/name", func(c *gin.Context) {
		c.String(200, "Yonatan Alebachew")
	})

	r.Run(":3000")
}
