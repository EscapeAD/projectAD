package routes

import (
	"escapead/projectAD/backend/api/controllers"
	"net/http"
)

var usersRoutes = []Route{
	Route{
		URI:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
	Route{
		URI:     "/users",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
	Route{
		URI:     "/users/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
	Route{
		URI:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	Route{
		URI:     "/users/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
	},
}
