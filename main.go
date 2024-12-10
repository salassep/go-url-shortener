package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salassep/go-url-shortener/handler"
	"github.com/salassep/go-url-shortener/store"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})

	r.POST("/create-short-url", handler.CreateShortUrl)
	r.GET("/:shortUrl", handler.RedirectShortUrl)

	store.InitializeStore()

	err := r.Run(":8080")

	if err != nil {
		panic(err)
	}
}
