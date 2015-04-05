package models

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Save the model
func Save(model Modeller) {
	if model.GetId() == "" {
		id := bson.NewObjectId()
		model.SetId(&id)
		log.Printf("Creating new object with id %s in %s", model.GetId(), model.Collection().FullName)

		model.Collection().Insert(model)
	} else {
		log.Printf("Updating object with id %s in %s", model.GetId(), model.Collection().FullName)

		model.Collection().Update(bson.M{"_id": model.GetId()}, model)
	}
}
