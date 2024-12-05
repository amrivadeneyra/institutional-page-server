package colegioapi

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetPathParams(ctx context.Context) map[string]string {
	pathParams := make(map[string]string)

	if rctx := chi.RouteContext(ctx); rctx != nil {
		for k := len(rctx.URLParams.Keys) - 1; k >= 0; k-- {
			pathParams[rctx.URLParams.Keys[k]] = rctx.URLParams.Values[k]
		}
	}

	return pathParams
}

func GetQueryParams(r *http.Request) map[string]string {
	queryParams := make(map[string]string)

	for key, value := range r.URL.Query() {
		queryParams[key] = value[0]
	}

	return queryParams
}
