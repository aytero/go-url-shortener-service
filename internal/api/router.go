package api

import (
	"github.com/gin-gonic/gin"
	"urlShortener/internal/app"
	"urlShortener/pkg/config"
)

type Api struct {
	app *app.App
	cfg *config.Config
}

func NewApi(app *app.App, cfg *config.Config) *Api {
	return &Api{
		app: app,
		cfg: cfg,
	}
}

func (a *Api) NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//router := gin.New()
	router.GET("/:shortUrl", a.GetOrigUrlByShort)
	router.POST("/", a.PostUrl)
	//router.GET("/swagger", httpSwagger)
	return router
}
