package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nakiscia/go-microservice-example/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId := req.URL.Query().Get("userId")

	user, err := services.GetUser(userId)
	if err != nil {
		// user not found
		resp.WriteHeader(int(err.StatusCode))
		appErr, _ := json.Marshal(err)
		resp.Write(appErr)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.WriteHeader(http.StatusOK)
	resp.Write(jsonValue)
}
