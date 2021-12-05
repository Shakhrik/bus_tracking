package api

import (
	"github.com/Shakhrik/inha/bus_tracking/config"
	"github.com/Shakhrik/inha/bus_tracking/pkg/logger"
	"github.com/Shakhrik/inha/bus_tracking/storage"

	v1 "github.com/Shakhrik/inha/bus_tracking/api/handlers/v1"

	_ "github.com/Shakhrik/inha/bus_tracking/api/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		apiV1.POST("/destination", handlerV1.DestinationCreate)
		apiV1.PUT("/destination/:id", handlerV1.DestinationUpdate)
		apiV1.GET("/destination/:id", handlerV1.DestinationGet)
		apiV1.GET("/destination", handlerV1.DestinationGetAll)
		apiV1.DELETE("/destination/:id", handlerV1.DestinationDelete)

		apiV1.POST("/bus-stop", handlerV1.BusStopCreate)
		apiV1.GET("/bus-stop/:destination_id", handlerV1.BusStopGetAll)
		apiV1.DELETE("/bus-stop/:id", handlerV1.BusStopDelete)

	}
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
