package domain

import (
	"fmt"
	"github.com/15matias15/golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Matias", LastName: "Pereira", Email: "my@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationErrors) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationErrors{
		Message:    fmt.Sprintf("users %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
