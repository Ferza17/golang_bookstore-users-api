package app

import (
	pingController "github.com/ferza17/golang_bookstore-users-api/controllers/ping"
	usersController "github.com/ferza17/golang_bookstore-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", pingController.Ping)

	// GET  USERS
	//router.GET("/users/", usersController.GetUsers)
	// GET USER
	router.GET("/users/:user_id", usersController.GetUser)
	// CREATE USER
	router.POST("/users", usersController.CreateUser)
	// UPDATE USER
	router.PUT("/users/:user_id", usersController.UpdateUser)
	// UPDATING USER (PARTIAL)
	router.PATCH("/users/:user_id", usersController.UpdateUser)
	// DELETE User
	router.DELETE("/users/:user_id", usersController.DeleteUser)
	// Search By status with query params ?status=""
	router.GET("internal/users/search", usersController.Search)
}
