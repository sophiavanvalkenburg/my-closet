package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sophiavanvalkenburg/my-closet/db"
	"github.com/sophiavanvalkenburg/my-closet/models"
)

type ItemController struct {
	db         *db.DB
	collection string
}

func NewItemController(db *db.DB, collection string) *ItemController {
	return &ItemController{db, collection}
}

func (itemc ItemController) ItemGet(w http.ResponseWriter, r *http.Request) {
	items := models.Items{}
	if err := itemc.db.FindItems(itemc.collection, 10, &items); err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(items); err != nil {
		fmt.Fprintln(w, "Error:", err)
	}
}

func (itemc ItemController) ItemGetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	item := models.Item{}
	if err := itemc.db.FindItemById(itemc.collection, itemId, &item); err != nil {
		switch err {
		case db.ErrInvalidId:
			w.WriteHeader(http.StatusBadRequest)
		case db.ErrCouldNotRetrieveDoc:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintln(w, "Error:", err)
		return
	}

	itemj, _ := json.Marshal(item)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", itemj)
}

func (itemc ItemController) ItemCreateOne(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}

	json.NewDecoder(r.Body).Decode(&item)
	itemc.db.InsertItem(itemc.collection, &item)

	itemj, _ := json.Marshal(item)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", itemj)
}

func (itemc ItemController) ItemUpdateOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	partialItem := models.Item{}
	json.NewDecoder(r.Body).Decode(&partialItem)
	if err := itemc.db.UpdateItemById(itemc.collection, itemId, &partialItem); err != nil {
		switch err {
		case db.ErrInvalidId:
		case db.ErrUnequalUpdateIds:
			w.WriteHeader(http.StatusBadRequest)
		case db.ErrCouldNotUpdateDoc:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintln(w, "Error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Item %s udpated.\n", itemId)
}

func (itemc ItemController) ItemDeleteOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	if err := itemc.db.RemoveById(itemc.collection, itemId); err != nil {
		switch err {
		case db.ErrInvalidId:
			w.WriteHeader(http.StatusBadRequest)
		case db.ErrCouldNotRemoveDoc:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintln(w, "Error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Item %s deleted.\n", itemId)

}
