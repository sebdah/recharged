package database

import "gopkg.in/mgo.v2"

func CreateCollectionBootNotificationLogs() {
	// Create the database
	database := GetDb()
	collection := mgo.Collection{
		Database: database,
		Name:     "bootNotificationLogs",
		FullName: database.Name + ".bootNotificationLogs",
	}
	collection.Create(&mgo.CollectionInfo{})
}

func GetCollectionBootNotificationLogs() *mgo.Collection {
	database := GetDb()
	return database.C("bootNotificationLogs")
}
