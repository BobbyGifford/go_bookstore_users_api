package users

import (
	"github.com/BobbyGifford/go_bookstore_users_api/domain/users"
	"github.com/BobbyGifford/go_bookstore_users_api/services"
	"github.com/BobbyGifford/go_bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

//	Input layer is controller package
func CreateUser(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
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

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}