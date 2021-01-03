package controllers

import (
	"github.com/ferza17/golang_bookstore-users-api/domain/users"
	"github.com/ferza17/golang_bookstore-users-api/services"
	"github.com/ferza17/golang_bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//func GetUsers(c *gin.Context) {}

func GetUser(c *gin.Context) {
	// GET PARAMS
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	// CHECK IF ERROR
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	// GET USER WITH SERVICES
	user, getErr := services.UserServices.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	// SEND RESPONSE
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public")=="true"))

}

//CreateUser
func  CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UserServices.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public")=="true"))
}

func  UpdateUser(c *gin.Context) {
	var user users.User
	// GET PARAMS
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	// GET JSON BODY
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// UPDATING USER
	user.ID = userId

	// IF PARTIAL UPDATE USING PATCH METHOD
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UserServices.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func  DeleteUser(c *gin.Context) {
	// GET PARAMS
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	// CHECK IF ERROR
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	if err := services.UserServices.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})
}

func  Search(c *gin.Context) {
	status := c.Query("status")
	result, err := services.UserServices.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public")=="true"))

}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UserServices.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
