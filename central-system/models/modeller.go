package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Modeller interface {
	GetId() bson.ObjectId
	Collection() *mgo.Collection
	SetId(id *bson.ObjectId)
}
