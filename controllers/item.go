package controllers

import (
    "fmt"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/sophiavanvalkenburg/my-closet/models"
    "github.com/sophiavanvalkenburg/my-closet/db"
)

type ItemController struct{
    db *db.DB 
    collection string
}

func NewItemController(db *db.DB, collection string) *ItemController {
    return &ItemController{db, collection}
}

func (itemc ItemController) ItemGet(w http.ResponseWriter, r *http.Request){
    items := models.Items{
        models.Item{Name: "Test1"},
        models.Item{Name: "Test2"},    
    } 

    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(items); err != nil {
        fmt.Fprintln(w, "Error:", err)
    }
}

func (itemc ItemController) ItemGetOne(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    itemId := vars["itemId"]
    item := models.Item{}
    if err := itemc.db.FindById(itemc.collection, itemId, &item); err != nil {
        switch err{
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
    fmt.Fprintf(w, "%s", itemj)
}

func (itemc ItemController) ItemCreateOne(w http.ResponseWriter, r *http.Request){ 
    item := models.Item{}

    json.NewDecoder(r.Body).Decode(&item)
    itemc.db.Insert(itemc.collection, &item)

    itemj, _ := json.Marshal(item)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%s\n", itemj)
}

func (itemc ItemController) ItemUpdateOne(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    itemId := vars["itemId"]
    fmt.Fprintln(w, "Update Item:", itemId)
}

func (itemc ItemController) ItemDeleteOne(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    itemId := vars["itemId"]
    fmt.Fprintln(w, "Delete Item:", itemId)
}