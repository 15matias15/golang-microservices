package services

import (
	"github.com/15matias15/golang-microservices/mvc/domain"
	"github.com/15matias15/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationErrors) {
	return domain.GetUser(userId)
}
