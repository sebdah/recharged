package models

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Ensure index
func EnsureIndexes(model Modeller) {
	log.Printf("Ensuring indexes for %s", model.Collection().FullName)
	for i := range model.Indexes() {
		model.Collection().EnsureIndex(*model.Indexes()[i])
	}
}

// Delete the model
func Delete(model Modeller) error {
	err := model.Collection().Remove(bson.M{"_id": model.GetId()})
	return err
}

// Save the model
func Save(model Modeller) (err error) {
	if model.GetId() == "" {
		id := bson.NewObjectId()
		model.SetId(&id)
		log.Printf(
			"Creating new object with id %s in %s",
			model.GetId().Hex(),
			model.Collection().FullName)

		err = model.Collection().Insert(model)
		if err != nil {
			log.Printf(
				"Error creating model %s (%s): %s",
				model.GetId().Hex(),
				model.Collection().FullName, err)
		}
	} else {
		log.Printf(
			"Updating object with id %s in %s",
			model.GetId().Hex(),
			model.Collection().FullName)

		err = model.Collection().Update(bson.M{"_id": model.GetId()}, model)
		if err != nil {
			log.Printf(
				"Error updating model %s (%s): %s",
				model.GetId().Hex(),
				model.Collection().FullName,
				err)
		}
	}

	return
}

// Update the model
func Update(model Modeller) (err error) {
	err = model.Collection().Update(bson.M{"_id": model.GetId()}, model)
	return
}
