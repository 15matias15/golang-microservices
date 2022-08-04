package app

import (
	"github.com/15matias15/golang-microservices/src/api/controllers/polo"
	"github.com/15matias15/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", repositories.CreateRepo)
}
