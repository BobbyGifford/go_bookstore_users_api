package app

import "github.com/BobbyGifford/go_bookstore_users_api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:user_id", controllers.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
}