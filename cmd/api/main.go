package cmd

import (
	// "fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Todo App is running",
			"status":  "success",
		})
	})

	router.Run(":3000")
}
