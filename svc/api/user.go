package colegioapi

import (
	"colegio/server/lib/jsonmodels"
	"colegio/server/svc/controllers"
	"encoding/json"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &jsonmodels.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createResult, validationErr, err := controllers.CreateUser(r.Context(), *user.ToModel())

	jsonResponse := jsonmodels.CreateResult{}
	jsonResponse.FillFromModel(createResult)

	handleResponse(w, r, jsonResponse, validationErr, err)
}

func VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := GetQueryParams(r)
	email := queryParams["email"]

	exists, err := controllers.VerifyEmail(r.Context(), email)

	handleResponse(w, r, exists, nil, err)
}
