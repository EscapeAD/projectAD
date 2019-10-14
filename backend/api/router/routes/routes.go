package routes

import (
	"escapead/projectAD/backend/api/middlewares"
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
	routes = append(routes, postsRoutes...)
	routes = append(routes, systemsRoutes...)
	return routes
}

// SetupRoutes - setup routes
func SetupRoutes(r chi.Router) chi.Router {
	for _, route := range Load() {
		r.MethodFunc(route.Method, route.URI, route.Handler)
	}
	return r
}

// SetupRoutesWithMiddleware - setup routes
func SetupRoutesWithMiddleware(r chi.Router) chi.Router {
	for _, route := range Load() {
		r.MethodFunc(route.Method, route.URI, middlewares.SetMiddlewareJSON(route.Handler))
	}
	return r
}
