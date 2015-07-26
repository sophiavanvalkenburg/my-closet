package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sophiavanvalkenburg/my-closet/controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(itemc *controllers.ItemController) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			controllers.Index,
		},
		Route{
			"ItemGet",
			"GET",
			"/api/v1/items.json",
			itemc.ItemsGetJson,
		},
		Route{
			"ItemGetOne",
			"GET",
			"/api/v1/items/{itemId}.json",
			itemc.ItemsGetOneJson,
		},
		Route{
			"ItemGet",
			"GET",
			"/items",
			itemc.ItemsGetHtml,
		},
		Route{
			"ItemGetOne",
			"GET",
			"/items/{itemId}",
			itemc.ItemsGetOneHtml,
		},
		Route{
			"ItemCreateOne",
			"PUT",
			"/api/v1/items.json",
			itemc.ItemsCreateOne,
		},
		Route{
			"ItemUpdateOne",
			"POST",
			"/api/v1/items/{itemId}.json",
			itemc.ItemsUpdateOne,
		},
		Route{
			"ItemDeleteOne",
			"DELETE",
			"/api/v1/items/{itemId}.json",
			itemc.ItemsDeleteOne,
		},
	}

	addRoutes(routes, router)
	return router
}

func addRoutes(routes Routes, router *mux.Router) {
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}
