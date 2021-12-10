package v1

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shakhrik/inha/bus_tracking/api/models"
	"github.com/Shakhrik/inha/bus_tracking/pkg/socket"
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

	go func() {
		message := `Bus operator with id = ` + strconv.Itoa(int(res.ID)) + ` has been created`
		socket.SocketClient(IP, PORT, message)
	}()

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
	limit, err := ParseQueryParam(c, "limit", "100")
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

//@Router /v1/login [post]
//@Summary Login user
//@Description API for login user
//@Tags user
//@Accept json
//@Produce json
//@Param user body models.UserLogin true "user"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) Login(c *gin.Context) {
	var userLogin models.UserLogin

	err := c.ShouldBind(&userLogin)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	res, err := h.storage.UserRepo().Login(userLogin)
	if err != nil {
		if err == sql.ErrNoRows {
			h.HandleErrorResponse(c, http.StatusUnauthorized, "incorrect login or password", nil)
			return
		}

		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 201, "user info", res)
}

//@Router /v1/user-buses/{user_id} [get]
//@Summary GetAll user buses
//@Description API for getting all user buses
//@Tags user
//@Accept json
//@Produce json
//@Param user_id path string true "user_id"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) GetAllUserBuses(c *gin.Context) {
	userID := c.Param("user_id")
	value, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.storage.UserRepo().GetUserBuses(int64(value))
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 200, "get all users successfully", res)
}
