package v1

import (
	"net/http"

	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/gin-gonic/gin"
)

func (h handlerV1) UserCreate(c *gin.Context) {
	var userCreate models.UserCreateRequest
	err := c.ShouldBindJSON(&userCreate)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	// res, err := h.storage.UserRepo()
}
