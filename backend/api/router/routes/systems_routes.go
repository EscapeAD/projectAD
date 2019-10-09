package routes

import (
	"escapead/projectAD/backend/api/controllers"
	"net/http"
)

var systemsRoutes = []Route{
	Route{
		URI:     "/version",
		Method:  http.MethodGet,
		Handler: controllers.GetVersion,
	},
	Route{
		URI:     "/Ping",
		Method:  http.MethodGet,
		Handler: controllers.Ping,
	},
}
