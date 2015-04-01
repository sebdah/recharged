package database

import "gopkg.in/mgo.v2"

func CreateCollectionIdTags() {
	database := GetDb()
	collection := mgo.Collection{
		Database: database,
		Name:     "idTags",
		FullName: database.Name + ".idTags",
	}
	collection.Create(&mgo.CollectionInfo{})
}
