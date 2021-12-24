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
	err := test.Run()
	if err != nil {
		log.Fatal(err)
	}
}
