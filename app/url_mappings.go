package app

import (
	"github.com/BobbyGifford/go_bookstore_users_api/controllers/ping"
	"github.com/BobbyGifford/go_bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
}