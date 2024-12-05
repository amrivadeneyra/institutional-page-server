package colegioapi

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Respond(w, r, fmt.Sprintf("Service is up and running"))
}
