package main

import (
	"github.com/aormcuw/ecom/cmd/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		panic(err)
	}

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.Run()
}
