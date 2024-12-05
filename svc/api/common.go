package colegioapi

import (
	"colegio/server/common/httpresponses"
	"colegio/server/common/utils"
	"net/http"
	"runtime/debug"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

func handleResponse(
	w http.ResponseWriter,
	r *http.Request,
	data interface{},
	validationErrors []*httpresponses.ValidationError,
	err error,
) {
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Respond(w, r, httpresponses.ErrorResponse{
			ErrorCode: err.Error(),
		})
		logrus.Errorf("%v\n\n\n%v", utils.GetErrorWithStack(err), string(debug.Stack()))
		return
	}
	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, &httpresponses.Response[*interface{}]{
		Data:             &data,
		ValidationErrors: validationErrors,
	})
}
