package colegioapi

import (
	"encoding/json"
	"net/http"
	"colegio/server/lib/jsonmodels"
	"colegio/server/svc/controllers"
)

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	refreshTokenData := &jsonmodels.RefreshToken{}

	if err := json.NewDecoder(r.Body).Decode(refreshTokenData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userToken, validationErr, err := controllers.RefreshToken(r.Context(), refreshTokenData.RefreshToken)

	jsonResponse := &jsonmodels.UserToken{}
	jsonResponse.FillFromModel(userToken)

	handleResponse(w, r, jsonResponse, validationErr, err)
}
