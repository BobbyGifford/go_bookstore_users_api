package users

import (
	"github.com/BobbyGifford/go_bookstore_users_api/domain/users"
	"github.com/BobbyGifford/go_bookstore_users_api/services"
	"github.com/BobbyGifford/go_bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	userId, userError := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userError != nil {
		err := errors.NewBadRequestError("user id is not a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}