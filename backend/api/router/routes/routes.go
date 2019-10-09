package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Route - Route Struct
type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// Load - Load routes
func Load() []Route {
	routes := usersRoutes
	routes = append(routes, systemsRoutes...)
	return routes
}

// SetupRoutes - setup routes
func SetupRoutes(r chi.Router) chi.Router {
	for _, route := range Load() {
		r.MethodFunc(route.Method, route.URI, route.Handler)
	}
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	return r
}
