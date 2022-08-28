package api

import (
	"github.com/aytero/go-url-shortener-service/internal/app"
	"github.com/aytero/go-url-shortener-service/pkg/config"
	"github.com/gin-gonic/gin"
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
	router.GET("/:shortUrl", a.GetFullUrlByShort)
	router.POST("/", a.PostUrl)
	return router
}

//server := &http.Server{
//	Addr:           ":8080",
//	Handler:        mux,
//	ReadTimeout:    10 * time.Second,
//	WriteTimeout:   10 * time.Second,
//	MaxHeaderBytes: 1 << 20,
//}
