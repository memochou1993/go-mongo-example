package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/memochou1993/movies-api/controllers"
)

// Route struct
type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("GET", "/movies", controllers.Index, nil)
	register("GET", "/movies/{id}", controllers.Show, nil)
	register("POST", "/movies", controllers.Store, nil)
	register("PUT", "/movies/{id}", controllers.Update, nil)
	register("DELETE", "/movies/{id}", controllers.Destroy, nil)
}

// NewRouter func
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Handler(route.Handler)

		if route.Middleware != nil {
			router.Use(route.Middleware)
		}
	}

	return router
}

func register(method string, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
