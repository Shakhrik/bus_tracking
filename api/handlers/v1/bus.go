package v1

import (
	"net/http"
	"strconv"

	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/Shakhrik/inha/bus_tracking/pkg/socket"
	"github.com/gin-gonic/gin"
)

const (
	IP   = "167.71.7.70"
	PORT = 9980
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

//@Router /v1/bus/{destination_id} [get]
//@Summary GetAll bus
//@Description API for getting all buses
//@Tags bus
//@Accept json
//@Produce json
//@Param destination_id path int true "destination_id"
//@Param limit query int false "limit"
//@Param page query int false "page"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) BusGetAll(c *gin.Context) {
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

	res, err := h.storage.BusRepo().GetAll(int32(value), limit, page)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 200, "get all buses successfully", res)
}

//@Router /v1/bus/reserve [post]
//@Summary Reserve bus
//@Description API for reserve bus
//@Tags bus
//@Accept json
//@Produce json
//@Param bus_reserve body models.BusReserve true "bus_reserve"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) BusReserve(c *gin.Context) {
	var busReserve models.BusReserve

	err := c.ShouldBind(&busReserve)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	res, err := h.storage.BusRepo().ReserveBus(busReserve.BusID)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 201, "bus reserved successfully", res)
}

//@Router /v1/bus/{id} [delete]
//@Summary Delete bus
//@Description API for deleting bus
//@Tags bus
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) BusDelete(c *gin.Context) {
	desID := c.Param("id")
	value, err := strconv.Atoi(desID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.storage.BusRepo().Delete(int64(value))
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 200, "bus deleted successfully", res)
}

//@Router /v1/bus [get]
//@Summary GetAll bus
//@Description API for getting all buses
//@Tags bus
//@Accept json
//@Produce json
//@Param limit query int false "limit"
//@Param page query int false "page"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) BriefBusGetAll(c *gin.Context) {
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

	res, err := h.storage.BusRepo().GetAllBuses(limit, page)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 200, "get all buses successfully", res)
}

//@Router /v1/change-status/{bus_id} [put]
//@Summary change status
//@Description API for changing status
//@Tags bus
//@Accept json
//@Produce json
//@Param bus_id path int true "bus_id"
//@Param bus_stop_id body models.ChangeStatus true "bus-stop-id"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) ChangeStatus(c *gin.Context) {
	busID := c.Param("bus_id")
	value, err := strconv.Atoi(busID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var changeStatus models.ChangeStatus
	changeStatus.BusID = int64(value)
	err = c.ShouldBind(&changeStatus)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	res, err := h.storage.BusRepo().ChangeStatus(changeStatus)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}
	socket.SocketClient(IP, PORT, strconv.Itoa(int(res.ID)))
	h.HandleSuccessResponse(c, 201, "bus created successfully", res)

}
