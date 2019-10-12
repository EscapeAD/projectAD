package router

import (
	"escapead/projectAD/backend/api/router/routes"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// var routes = flag.Bool("routes", false, "Generate router documentation")

// NEW - new Router instance
func NEW() chi.Router {
	// TODO use flag for documention generation
	// flag.Parse()
	r := chi.NewRouter()

	// LOAD MIDDLE WARES
	// Injects a request ID into the context of each request
	r.Use(middleware.RequestID)

	//Sets a http.Request's RemoteAddr to either X-Forwarded-For or X-Real-IP
	r.Use(middleware.RealIP)

	//Logs the start and end of each request with the elapsed processing time
	// r.Use(middleware.Logger)

	//Gracefully absorb panics and prints the stack trace
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	return routes.SetupRoutesWithMiddleware(r)
}
