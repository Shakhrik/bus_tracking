package v1

import (
	"net/http"
	"strconv"

	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/Shakhrik/inha/bus_tracking/config"
	"github.com/Shakhrik/inha/bus_tracking/pkg/logger"
	"github.com/Shakhrik/inha/bus_tracking/storage"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log     logger.Logger
	cfg     *config.Config
	storage storage.StorageI
}

type HandlerV1Options struct {
	Log     logger.Logger
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		log:     options.Log,
		cfg:     options.Cfg,
		storage: options.Storage,
	}
}

func ParseQueryParam(c *gin.Context, key string, defaultValue string) (int32, error) {
	valueStr := c.DefaultQuery(key, defaultValue)
	if valueStr == "" {
		return 0, nil
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 0, err
	}

	return int32(value), nil
}

func (h *handlerV1) HandleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	h.log.Error(message, logger.Int("code", code), logger.Any("error", err))
	c.JSON(code, models.ErrorModel{
		Code:    code,
		Message: message,
		Error:   err,
	})
}
func (h *handlerV1) HandleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, models.SuccessModel{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
