package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	test := gin.Default()

	test.GET("/healthcheck", healthCheck)
	v1 := test.Group("v1")
	{
		v1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"testing": "testin1234",
			})
		})
		v1.POST("/testpost", testPost)
	}
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

func testPost(ctx *gin.Context) {
	name := ctx.PostForm("name")
	value := ctx.PostForm("value")

	if len(name) > 0 && len(value) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("My name is %s and you have %s value to me,", name, value),
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "you've not provided the appropriate name and value values, cheers.",
		})
	}
}
