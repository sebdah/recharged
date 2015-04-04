package database

import "gopkg.in/mgo.v2"

var collectionName string = "idTags"

func CreateCollectionIdTags() {
	database := GetDb()
	collection := mgo.Collection{
		Database: database,
		Name:     collectionName,
		FullName: database.Name + "." + collectionName,
	}
	collection.Create(&mgo.CollectionInfo{})
}

func GetCollectionIdTags() *mgo.Collection {
	database := GetDb()
	return database.C(collectionName)
}
