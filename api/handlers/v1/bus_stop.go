package v1

import (
	"net/http"
	"strconv"

	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/gin-gonic/gin"
)

//@Router /v1/bus-stop [post]
//@Summary Create bus-stop
//@Description API for create bus-stop
//@Tags bus-stop
//@Accept json
//@Produce json
//@Param employee body models.BusStopCreate true "bus-stop"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) BusStopCreate(c *gin.Context) {
	var busStopCreate models.BusStopCreate

	err := c.ShouldBind(&busStopCreate)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	res, err := h.storage.BusStopRepo().Create(busStopCreate)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 201, "bus-stop created successfully", res)
}

//@Router /v1/bus-stop/{destination_id} [get]
//@Summary GetAll bus-stop
//@Description API for getting all bus-stopes
//@Tags bus-stop
//@Accept json
//@Produce json
//@Param destination_id path int true "destination_id"
//@Param limit query int false "limit"
//@Param page query int false "page"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) BusStopGetAll(c *gin.Context) {
	desID := c.Param("destination_id")
	value, err := strconv.Atoi(desID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	limit, err := ParseQueryParam(c, "limit", "10")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	page, err := ParseQueryParam(c, "page", "1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.storage.BusStopRepo().GetAll(int32(value), limit, page)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 200, "get all bus-stopes successfully", res)
}

//@Router /v1/bus-stop/{id} [delete]
//@Summary Delete bus-stop
//@Description API for deleting bus-stop
//@Tags bus-stop
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) BusStopDelete(c *gin.Context) {
	desID := c.Param("id")
	value, err := strconv.Atoi(desID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.storage.BusStopRepo().Delete(int64(value))
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 200, "bus-stop deleted successfully", res)
}
