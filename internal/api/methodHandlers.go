package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlShortener/pkg/model"
)

//curl -X POST http://localhost:8080/ -H 'Content-Type: application/json' -d '{"short_url":"abb","full_url":"http//:example.com/dddddddfs"}'
//curl http://localhost:8080/abb

func (a *Api) PostUrl(c *gin.Context) {
	var request model.PostUrl

	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Internal server error"})
		return
	}
	shortUrl, err := a.app.ShortenUrl(request.FullUrl)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Internal server error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, shortUrl)
}

func (a *Api) GetOrigUrlByShort(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	if shortUrl == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Internal server error"})
		return
	}
	fullUrl, err := a.app.GetFullUrl(shortUrl)
	if fullUrl == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "URL not found"})
	} else if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Internal server error"})
	}
	c.IndentedJSON(http.StatusOK, fullUrl)

}
