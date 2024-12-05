package colegioapi

import (
	"colegio/server/svc/middleware"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Authorization", "X-Auth-Token", "X-Api-Key", "Access-Control-Allow-Origin"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/api/v1", func(r chi.Router) {
		r.Use(apiVersionCtx("v1"))
		r.Get("/healthcheck", HealthcheckHandler)

		// User
		r.Route("/user", func(r chi.Router) {
			//r.Get("/verify/email", VerifyEmailHandler)
			r.Group(func(r chi.Router) {
				r.Use(middleware.RequiresAuth)
				r.Post("/", CreateUserHandler)
			})
		})

		// Login
		r.Route("/login", func(r chi.Router) {
			//r.Post("/", LoginHandler)
		})

		//Refresh
		r.Route("/refresh", func(r chi.Router) {
			//r.Post("/", RefreshTokenHandler)
		})

		r.Route("/logout", func(r chi.Router) {
			r.Use(middleware.RequiresAuth)
			//r.Post("/", LogoutHandler)
		})

		// Task
		r.Route("/task", func(r chi.Router) {
			//r.Use(middleware.RequiresAuth)
			//r.Get("/{id}", GetTaskHandler)
			//r.Post("/", CreateTaskHandler)
			//r.Put("/{id}", UpdateTaskHandler)
			//r.Put("/dueDate/{id}", UpdateTaskDueDateHandler)
			//r.Delete("/{id}", DeleteTaskHandler)
			//r.Get("/currentUser", GetTasksByCurrentUserHandler)
			//r.Get("/draggable/{id}", GetTasksForDraggableHandler)
			//r.Get("/", GetTasksHandler)
		})

		r.Route("/user-settings", func(r chi.Router) {
			//r.Use(middleware.RequiresAuth)
			//r.Get("/", GetUserSettingsHandler)
			//r.Post("/", CreateUserSettingsHandler)
			//r.Put("/{id}", UpdateUserSettingsHandler)
		})
	})

	return router
}

func getHandler(requestPath string, ID string) http.Handler {
	switch requestPath {
	case "email":
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func postHandler(requestPath string, ID string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func putHandler(requestPath string, ID string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func deleteHandler(requestPath string, ID string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

type APIContextKey string

func apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), APIContextKey("api.version"), version))
			next.ServeHTTP(w, r)
		})
	}
}
