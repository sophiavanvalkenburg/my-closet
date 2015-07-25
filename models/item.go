package models

import "gopkg.in/mgo.v2/bson"

type Item struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	Type          string        `json:"type" bson: "name"`
	Brand         string        `json:"brand" bson:"brand"`
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

func (item *Item) CreatePartialUpdateMap() map[string]interface{} {
	updateMap := map[string]interface{}{}
	if item.Type != "" {
		updateMap["type"] = item.Type
	}
	if item.Brand != "" {
		updateMap["brand"] = item.Brand
	}
	if item.Name != "" {
		updateMap["name"] = item.Name
	}
	if item.Color != "" {
		updateMap["color"] = item.Color
	}
	if item.Size != "" {
		updateMap["size"] = item.Size
	}
	if item.Collection != "" {
		updateMap["collection"] = item.Collection
	}
	if item.Tags != nil {
		updateMap["tags"] = item.Tags
	}
	if item.ImageLink != "" {
		updateMap["image_link"] = item.ImageLink
	}
	if item.ReferenceLink != "" {
		updateMap["reference_link"] = item.ReferenceLink
	}
	if item.SaleLink != "" {
		updateMap["sale_link"] = item.SaleLink
	}
	return updateMap
}
