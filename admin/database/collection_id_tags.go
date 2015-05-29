package database

import "gopkg.in/mgo.v2"

func CreateCollectionIdTags() {
	// Create the database
	database := GetDb()
	collection := mgo.Collection{
		Database: database,
		Name:     "idTags",
		FullName: database.Name + ".idTags",
	}
	collection.Create(&mgo.CollectionInfo{})
}

func GetCollectionIdTags() *mgo.Collection {
	database := GetDb()
	return database.C("idTags")
}
