package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	app.GET("/hello", func(ctx *gin.Context){
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.String(200, "Hello")
	})

	app.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}