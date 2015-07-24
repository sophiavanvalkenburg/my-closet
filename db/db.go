package db

import (
	"errors"
	"fmt"

	"github.com/sophiavanvalkenburg/my-closet/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrInvalidId           = errors.New("Invalid Object Id")
	ErrCouldNotRetrieveDoc = errors.New("Could Not Retrieve Object")
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

func (theDB *DB) Insert(collection string, docPtr models.UniqueModel) {
	fmt.Println("Trying to insert into DB...")
	docPtr.SetUniqueId()
	err := theDB.session.DB(theDB.name).C(collection).Insert(docPtr)
	if err != nil {
		printDBError(err)
	}
}

func (theDB *DB) FindById(collection string, id string, docPtr models.UniqueModel) error {
	if !bson.IsObjectIdHex(id) {
		return ErrInvalidId
	}
	oid := bson.ObjectIdHex(id)

	if err := theDB.session.DB(theDB.name).C(collection).FindId(oid).One(docPtr); err != nil {
		printDBError(err)
		return ErrCouldNotRetrieveDoc
	}

	return nil
}

func printDBError(err error) {
	fmt.Println("DB ERROR:", err)
}
