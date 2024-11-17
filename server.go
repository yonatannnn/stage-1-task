package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/name", func(c *gin.Context) {
		c.String(200, "Yonatan Alebachew")
	})

	r.GET("/hobby", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hobby": "Playing Soccer",
		})
	})


	r.Run(":3000")
}
