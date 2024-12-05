package middleware

import (
	"colegio/server/lib/auth"
	"colegio/server/lib/auth/selfauthorizer"
	"net/http"
	"strings"
)

func RequiresAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		authorizer := selfauthorizer.NewSelfAuthorizer()
		parsedToken, err := authorizer.Validate(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx, err := auth.DecorateContext(r.Context(), parsedToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
