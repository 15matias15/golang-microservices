package services

import (
	"github.com/15matias15/golang-microservices/mvc/domain"
	"github.com/15matias15/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDaoMock     usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationErrors)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct {
}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationErrors) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDB(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationErrors) {
		return nil, &utils.ApplicationErrors{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 does not exists",
		}
	}
	user, err := UsersServices.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exists", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationErrors) {
		return &domain.User{
			Id: 123,
		}, nil
	}
	user, err := UsersServices.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
}
