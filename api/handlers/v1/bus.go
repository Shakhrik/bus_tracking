package v1

import (
	"net/http"

	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/gin-gonic/gin"
)

//@Router /v1/bus [post]
//@Summary Create bus
//@Description API for create bus
//@Tags bus
//@Accept json
//@Produce json
//@Param employee body models.BusCreate true "bus"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) BusCreate(c *gin.Context) {
	var busCreate models.BusCreate

	err := c.ShouldBind(&busCreate)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	res, err := h.storage.BusRepo().Create(busCreate)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 201, "bus created successfully", res)
}
