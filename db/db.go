package db

import (
	//"encoding/json"
	"errors"
	"fmt"

	"github.com/sophiavanvalkenburg/my-closet/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrInvalidId            = errors.New("Invalid Object Id")
	ErrCouldNotRetrieveDoc  = errors.New("Could Not Retrieve Object")
	ErrCouldNotRetrieveDocs = errors.New("Could Not Retrieve Objects")
	ErrCouldNotRemoveDoc    = errors.New("Could Not Remove Object")
	ErrCouldNotUpdateDoc    = errors.New("Could Not Update Object")
	ErrUnequalUpdateIds     = errors.New("Update Doc Ids Do Not Match")
)

type DB struct {
	session *mgo.Session
	name    string
}

func NewDB(name string) *DB {
	fmt.Println("Initiating DB connection")
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		printDBError(err)
		return nil
	}
	return &DB{s, name}
}

func (theDB *DB) InsertItem(collection string, item *models.Item) {
	item.SetUniqueId()
	err := theDB.session.DB(theDB.name).C(collection).Insert(item)
	if err != nil {
		printDBError(err)
	}
}

func (theDB *DB) FindItems(collection string, limit int, res *models.Items) error {
	iter := theDB.session.DB(theDB.name).C(collection).Find(nil).Limit(limit).Iter()
	if err := iter.All(res); err != nil {
		printDBError(err)
		return ErrCouldNotRetrieveDocs
	}

	return nil
}
func (theDB *DB) FindItemById(collection string, id string, item *models.Item) error {
	if !bson.IsObjectIdHex(id) {
		return ErrInvalidId
	}
	oid := bson.ObjectIdHex(id)

	if err := theDB.session.DB(theDB.name).C(collection).FindId(oid).One(item); err != nil {
		printDBError(err)
		return ErrCouldNotRetrieveDoc
	}

	return nil
}

func (theDB *DB) UpdateItemById(collection string, id string, partialItem *models.Item) error {
	if !bson.IsObjectIdHex(id) {
		return ErrInvalidId
	}
	oid := bson.ObjectIdHex(id)

	if partialItem.Id != "" && partialItem.Id != oid {
		return ErrUnequalUpdateIds
	}
	updateMap := map[string]interface{}{
		"$set": partialItem.CreatePartialUpdateMap(),
	}
	if err := theDB.session.DB(theDB.name).C(collection).UpdateId(oid, updateMap); err != nil {
		printDBError(err)
		return ErrCouldNotUpdateDoc
	}

	return nil
}
func (theDB *DB) RemoveById(collection string, id string) error {
	if !bson.IsObjectIdHex(id) {
		return ErrInvalidId
	}
	oid := bson.ObjectIdHex(id)

	if err := theDB.session.DB(theDB.name).C(collection).RemoveId(oid); err != nil {
		printDBError(err)
		return ErrCouldNotRemoveDoc
	}

	return nil
}

func printDBError(err error) {
	fmt.Println("DB ERROR:", err)
}
