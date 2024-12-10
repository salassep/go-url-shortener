package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salassep/go-url-shortener/shortener"
	"github.com/salassep/go-url-shortener/store"
	"github.com/salassep/go-url-shortener/util"
)

type CreateShortUrlRequest struct {
	OriginalUrl string `json:"originalUrl" binding:"required"`
	UserId      string `json:"userId" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var createShortUrlRequest CreateShortUrlRequest
	if err := c.ShouldBindJSON(&createShortUrlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(createShortUrlRequest.OriginalUrl, createShortUrlRequest.UserId)
	store.SaveUrlMapping(shortUrl, createShortUrlRequest.OriginalUrl)

	host := util.GetRootUrl(c)

	c.JSON(200, gin.H{
		"message":  "Create short url success",
		"shortUrl": host + shortUrl,
	})
}

func RedirectShortUrl(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)

	if initialUrl == "" {
		c.JSON(404, gin.H{"error": "Short url not found"})
		return
	}

	c.Redirect(302, initialUrl)
}
