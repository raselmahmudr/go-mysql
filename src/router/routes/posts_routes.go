package routes

import (
	"hello/src/api/controllers"
	"net/http"
)


var postsRoutes = []Route{
	Route{
		Uri: "/api/posts",
		Method: http.MethodGet,
		Handler: controllers.GetPost,
	},
	Route{
		Uri: "/api/post",
		Method: http.MethodPost,
		Handler: controllers.CreatePost,
	},
}

