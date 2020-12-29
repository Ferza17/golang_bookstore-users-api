package controllers

import (
	"net/http"
	"strconv"

	"github.com/ferza17/golang_bookstore-users-api/domain/users"
	"github.com/ferza17/golang_bookstore-users-api/services"
	"github.com/ferza17/golang_bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {}

func GetUser(c *gin.Context) {
	// GET PARAMS
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	// CHECK IF ERROR
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	// GET USER WITH SERBICES
	user, getErr := services.GetUser(userId)
	if getErr != nil{
		c.JSON(getErr.Status, getErr)
		return
	}
	// SEND RESPONSE
	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func UpdateUser(c *gin.Context)  {
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


	result,err := services.UpdateUser(isPartial,user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}