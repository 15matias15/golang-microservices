package services

import (
	"github.com/15matias15/golang-microservices/mvc/domain"
	"github.com/15matias15/golang-microservices/mvc/utils"
)

type usersService struct{}

var (
	UsersServices usersService
)

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationErrors) {
	return domain.UserDao.GetUser(userId)
}
