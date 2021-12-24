package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	test := gin.Default()
	test.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"testing": "testin1234",
		})
	})
	test.GET("/healthcheck", healthCheck)
	err := test.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"serviceStatus": "OK",
		"hostname":      "localhost",
	})
}
