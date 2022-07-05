package app

import (
	"github.com/15matias15/golang-microservices/mvc/controllers"
	"net/http"
)

func StartApplication() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
