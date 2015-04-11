package database

import "gopkg.in/mgo.v2"

var collectionName string = "chargePoints"

func CreateCollectionChargePoints() {
	database := GetDb()
	collection := mgo.Collection{
		Database: database,
		Name:     collectionName,
		FullName: database.Name + "." + collectionName,
	}
	collection.Create(&mgo.CollectionInfo{})
}

func GetCollectionChargePoints() {
	database := GetDb()
	return database.C(collectionName)
}
