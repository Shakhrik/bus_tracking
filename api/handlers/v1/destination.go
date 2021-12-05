package v1

import (
	"net/http"
	"strconv"

	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/gin-gonic/gin"
)

//@Router /v1/destination [post]
//@Summary Create destination
//@Description API for create destination
//@Tags destination
//@Accept json
//@Produce json
//@Param employee body models.DestinationCreate true "destination"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) DestinationCreate(c *gin.Context) {
	var destinationCreate models.DestinationCreate

	err := c.ShouldBind(&destinationCreate)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	res, err := h.storage.DestinationRepo().Create(destinationCreate)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 201, "destination created successfully", res)
}

//@Router /v1/destination/{id} [put]
//@Summary Update destination
//@Description API for updating destination
//@Tags destination
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Param employee body models.DestinationUpdate true "destination"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) DestinationUpdate(c *gin.Context) {
	var destinationUpdate models.DestinationUpdate

	err := c.ShouldBind(&destinationUpdate)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	desID := c.Param("id")
	value, err := strconv.Atoi(desID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	destinationUpdate.ID = int64(value)
	res, err := h.storage.DestinationRepo().Update(destinationUpdate)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 200, "destination updated successfully", res)
}

//@Router /v1/destination/{id} [get]
//@Summary Get destination
//@Description API for getting destination
//@Tags destination
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) DestinationGet(c *gin.Context) {
	desID := c.Param("id")
	value, err := strconv.Atoi(desID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.storage.DestinationRepo().Get(int64(value))
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 200, "get successfully", res)
}

//@Router /v1/destination [get]
//@Summary GetAll destinations
//@Description API for getting all destinations
//@Tags destination
//@Accept json
//@Produce json
//@Param limit query int false "limit"
//@Param page query int false "page"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) DestinationGetAll(c *gin.Context) {
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

	res, err := h.storage.DestinationRepo().GetAll(limit, page)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 200, "get all successfully", res)
}

//@Router /v1/destination/{id} [delete]
//@Summary Delete destination
//@Description API for deleting destination
//@Tags destination
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) DestinationDelete(c *gin.Context) {
	desID := c.Param("id")
	value, err := strconv.Atoi(desID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.storage.DestinationRepo().Delete(int32(value))
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 200, "destination deleted successfully", res)
}
