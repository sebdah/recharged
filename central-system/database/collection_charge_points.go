package database

import "gopkg.in/mgo.v2"

func CreateCollectionChargePoints() {
	database := GetDb()
	collection := mgo.Collection{
		Database: database,
		Name:     "chargePoints",
		FullName: database.Name + ".chargePoints",
	}
	collection.Create(&mgo.CollectionInfo{})
}

func GetCollectionChargePoints() *mgo.Collection {
	database := GetDb()
	return database.C("chargePoints")
}
