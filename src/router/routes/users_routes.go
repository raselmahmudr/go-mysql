

package routes

import (
	"hello/src/api/controllers"
	"net/http"
)

var usersRoutes = []Route{
	Route{
		Uri: "/users",
		Method: http.MethodGet,
		Handler: controllers.GetUsers,
	},
	Route{
		Uri: "/user",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
	},
}