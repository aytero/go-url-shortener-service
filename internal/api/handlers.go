package api

import (
	"github.com/aytero/go-url-shortener-service/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func (a *Api) PostUrl(c *gin.Context) {
	var request model.PostUrl

	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	_, err := url.ParseRequestURI(request.FullUrl)
	if err != nil {
		a.cfg.ErrorLog.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "Invalid URL",
		})
		return
	}
	shortUrl, err := a.app.ShortenUrl(request.FullUrl)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Internal server error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message":   "Success",
		"short_url": shortUrl,
	})
}

func (a *Api) GetFullUrlByShort(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	if shortUrl == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	fullUrl, err := a.app.GetFullUrl(shortUrl)
	if fullUrl == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "URL not found"})
	} else if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Internal server error"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"full_url": fullUrl})

}
