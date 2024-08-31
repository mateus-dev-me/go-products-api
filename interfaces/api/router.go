package api

import (
	"go-products/interfaces/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(c *controllers.ProductController) *gin.Engine {
	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	group := server.Group("/api/v1")

	group.GET("/products", c.GetAllHandler)
	group.GET("/products/:id", c.GetByIDHandler)
	group.POST("/products", c.SaveHandler)
	group.PUT("/products/:id", c.UpdateHandler)
	group.DELETE("/products/:id", c.DeleteHandler)

	return server
}
