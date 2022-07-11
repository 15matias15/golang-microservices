package app

import "github.com/15matias15/golang-microservices/mvc/controllers"

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)

}
