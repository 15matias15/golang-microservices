package controllers

import (
	"github.com/15matias15/golang-microservices/mvc/services"
	"github.com/15matias15/golang-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationErrors{
			Message:    "user_id must be of type int64",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		c.JSON(http.StatusBadRequest, apiErr)
		return
	}

	user, apiErr := services.UsersServices.GetUser(userId)
	if apiErr != nil {
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
