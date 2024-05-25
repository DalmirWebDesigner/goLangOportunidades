package router

import (
	"github.com/gin-gonic/gin"
)

func Inicialize() {
		// Inicializa o Router com config default
		router := gin.Default()

		// Define uma rota
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		// Rodando a api
		router.Run(":8080") // listen and serve on 0.0.0.0:8080
	}