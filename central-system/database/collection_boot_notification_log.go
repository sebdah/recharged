package database

import "gopkg.in/mgo.v2"

var collectionName string = "bootNotificationLog"

func CreateCollectionBootNotificationLog() {
	// Create the database
	database := GetDb()
	collection := mgo.Collection{
		Database: database,
		Name:     "bootNotificationLog",
		FullName: database.Name + ".bootNotificationLog",
	}
	collection.Create(&mgo.CollectionInfo{})
}

func GetCollectionBootNotificationLog() *mgo.Collection {
	database := GetDb()
	return database.C("bootNotificationLog")
}
