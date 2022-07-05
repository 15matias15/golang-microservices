package controllers

import (
	"encoding/json"
	"github.com/15matias15/golang-microservices/mvc/services"
	"github.com/15matias15/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationErrors{
			Message:    "user_id must be of type int64",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonErrorValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonErrorValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
