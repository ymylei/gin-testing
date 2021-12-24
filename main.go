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

type Test struct {
	Name  string `form:"name" json:"name" binding:"required"`
	Value string `form:"value" json:"value" binding:"required"`
}

func testPost(ctx *gin.Context) {
	var form Test
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":        "you've goofed the form, cheers.",
			"error_return": err.Error(),
		})
		return
	}

	if len(form.Name) > 0 && len(form.Value) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("My name is %s and you have %s value to me.", form.Name, form.Value),
		})
		return
	}

	ctx.JSON(http.StatusTeapot, gin.H{
		"error": "you've done it now mate.",
	})
}
