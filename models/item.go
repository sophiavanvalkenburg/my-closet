package models

import "gopkg.in/mgo.v2/bson"

type Item struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	Type          string        `json:"type" bson: "name"`
	Brand         string        `json:"completed" bson:"completed"`
	Name          string        `json:"name" bson:"name"`
	Color         string        `json:"color" bson:"color"`
	Size          string        `json:"size" bson:"size"`
	Collection    string        `json:"collection" bson:"collection`
	Tags          []string      `json:"tags" bson:"tags"`
	ImageLink     string        `json:"image_link" bson:"image_link"`
	ReferenceLink string        `json:"reference_link" bson:"reference_link"`
	SaleLink      string        `json:"sale_link" bson:"sale_link"`
}

type Items []Item

func (item *Item) SetUniqueId() {
	item.Id = bson.NewObjectId()
}
