package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetRootUrl(c *gin.Context) string {
	scheme := "http"

	if c.Request.TLS != nil {
		scheme = "https"
	}

	host := c.Request.Host

	return fmt.Sprintf("%s://%s/", scheme, host)
}
