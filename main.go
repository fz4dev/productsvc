package main

import (
	"github.com/fzuchows/productsvc/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/product/:id", func(c *gin.Context) {
		product, err := database.GetProduct(c.Param("id"))
		if err != nil {
			c.String(404, err.Error())
			return
		}
		c.JSON(200, product)
	})

	r.Run(":8900")
}
