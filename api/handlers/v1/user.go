package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/gin-gonic/gin"
)

//@Router /v1/user [post]
//@Summary Create user
//@Description API for create user
//@Tags user
//@Accept json
//@Produce json
//@Param user body models.UserCreateRequest true "user"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) UserCreate(c *gin.Context) {
	var userCreate models.UserCreateRequest

	err := c.ShouldBind(&userCreate)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	res, err := h.storage.UserRepo().Create(userCreate)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 201, "user created successfully", res)
}

//@Router /v1/user/{id} [delete]
//@Summary Delete user
//@Description API for deleting user
//@Tags user
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) UserDelete(c *gin.Context) {
	desID := c.Param("id")
	value, err := strconv.Atoi(desID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.storage.UserRepo().Delete(int32(value))
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 200, "user deleted successfully", res)
}

//@Router /v1/user [get]
//@Summary GetAll users
//@Description API for getting all users
//@Tags user
//@Accept json
//@Produce json
//@Param limit query int false "limit"
//@Param alias query string false "alias"
//@Param page query int false "page"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) UserGetAll(c *gin.Context) {
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

	alias := c.DefaultQuery("alias", "")

	res, err := h.storage.UserRepo().GetAll(alias, limit, page)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 200, "get all users successfully", res)
}

func (h handlerV1) Socket(c *gin.Context) {
	fmt.Println("keldiii", c.Request.Body)

}
