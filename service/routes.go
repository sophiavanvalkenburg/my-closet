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
            "/items",
            itemc.ItemGet,
        },
        Route{
            "ItemGetOne",
            "GET",
            "/items/{itemId}",
            itemc.ItemGetOne,
        },
        Route{
            "ItemCreateOne",
            "PUT",
            "/items",
            itemc.ItemCreateOne,
        },
        Route{
            "ItemUpdateOne",
            "POST",
            "/items/{itemId}",
            itemc.ItemUpdateOne,
        },
        Route{
            "ItemDeleteOne",
            "DELETE",
            "/items/{itemId}",
            itemc.ItemDeleteOne,
        },
    }

    addRoutes(routes, router)
    return router
}

func addRoutes( routes Routes, router *mux.Router){
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
