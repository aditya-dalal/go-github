package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/go-github/lib"
)

type route struct {
	Name string
	Method string
	Pattern string
	Handler func(lib.AppContext, http.ResponseWriter, *http.Request) (int, error)
}

func CreateRouter(c lib.AppContext) *mux.Router {
	routes := map[string][]route{
		"/events": events,
		"/actors": actors,
		"/erase": erase,
	}

	router := mux.NewRouter().StrictSlash(true)
	for prefix, pathRoutes := range routes {
		pathRouter := router.PathPrefix(prefix).Subrouter()
		for _, route := range pathRoutes {
			pathRouter.
				Methods(route.Method).
				Name(route.Name).
				Path(route.Pattern).
				Handler(lib.AppHandler{c, route.Handler})
		}
	}
	return router
}
