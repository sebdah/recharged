package models

type Modeller interface {
	Save() *mgo.Document
}
