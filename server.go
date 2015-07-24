/*******
server.go - starts a simple web server

*******/

package main

import (
	"log"
	"net/http"
	//"fmt"

	"github.com/sophiavanvalkenburg/my-closet/controllers"
	"github.com/sophiavanvalkenburg/my-closet/db"
	"github.com/sophiavanvalkenburg/my-closet/service"
)

func main() {

	theDB := db.NewDB("my_closet_test")
	if theDB != nil {
		itemc := controllers.NewItemController(theDB, "items")
		router := service.NewRouter(itemc)
		log.Fatal(http.ListenAndServe(":6060", router))
	}
}
