package domain

import (
	"fmt"
	"github.com/15matias15/golang-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Matias", LastName: "Pereira", Email: "my@gmail.com"},
	}
	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationErrors)
}

type userDao struct {
}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationErrors) {
	log.Println("we're accessing the database ")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationErrors{
		Message:    fmt.Sprintf("users %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
