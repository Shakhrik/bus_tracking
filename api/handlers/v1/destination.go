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
