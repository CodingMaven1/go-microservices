package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CodingMaven1/go-microservices/mvc/services"
	"github.com/CodingMaven1/go-microservices/mvc/utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "User ID must be number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonresponse, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write([]byte(jsonresponse))
		return
	}

	user, apiError := services.GetUser(userId)
	if apiError != nil {
		jsonresponse, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write([]byte(jsonresponse))
		return
	}

	jsonresponse, err := json.Marshal(user)
	if err != nil {
		return
	}
	res.Write(jsonresponse)
}
