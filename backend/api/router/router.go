package router

import (
	"escapead/projectAD/backend/api/router/routes"

	"github.com/go-chi/chi"
)

// var routes = flag.Bool("routes", false, "Generate router documentation")

// NEW - new Router instance
func NEW() chi.Router {
	// TODO use flag for documention generation
	// flag.Parse()
	r := chi.NewRouter()
	return routes.SetupRoutes(r)
}
