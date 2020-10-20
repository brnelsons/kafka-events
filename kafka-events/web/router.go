package web

import "github.com/gorilla/mux"

func NewRouter(rootPath string, routes Routes) *mux.Router {
	// When StrictSlash == true, if the route path is "/path/", accessing "/path" will perform a redirect to the former and vice versa.
	router := mux.NewRouter().StrictSlash(true)
	router.Use(Logger)
	sub := router.PathPrefix(rootPath).Subrouter()

	for _, route := range routes {
		sub.
			HandleFunc(route.Pattern, route.HandlerFunc).
			Name(route.Name).
			Methods(route.Method)
	}

	return router
}
