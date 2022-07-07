package services

import (
	"github.com/15matias15/golang-microservices/mvc/domain"
	"github.com/15matias15/golang-microservices/mvc/utils"
	"net/http"
)

type itemsService struct {
}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationErrors) {
	return nil, &utils.ApplicationErrors{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
