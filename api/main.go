package api

import (
	"github.com/Shakhrik/bus_tracking/api/storage"

	v1 "github.com/Shakhrik/bus_tracking/api/api/handlers/v1"

	"github.com/Shakhrik/bus_tracking/api/config"
	"github.com/Shakhrik/bus_tracking/api/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterOptions struct {
	Log     logger.Logger
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	cfg := cors.DefaultConfig()

	cfg.AllowHeaders = append(cfg.AllowHeaders, "*")
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true

	router.Use(cors.New(cfg))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:     opt.Log,
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	apiV1 := router.Group("/v1")

	// apiV1.Use(handlerV1.AuthMiddleware())
	{
		apiV1.GET("", handlerV1.HandleErrorResponse)
	}

	return router
}
