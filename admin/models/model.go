package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Ensure index
func EnsureIndexes(model Modeller) {
	log.Debug("Ensuring indexes for %s", model.Collection().FullName)
	for i := range model.Indexes() {
		model.Collection().EnsureIndex(*model.Indexes()[i])
	}
}

// Delete the model
func Delete(model Modeller) error {
	err := model.Collection().Remove(bson.M{"_id": model.GetId()})
	return err
}

// Drop collection
func DropCollection(model Modeller) error {
	log.Info("Dropping collection '%s'\n", model.Collection().Name)
	err := model.Collection().DropCollection()
	return err
}

// Save the model
func Save(model Modeller) (err error) {
	if model.GetId() == "" {
		id := bson.NewObjectId()
		model.SetId(&id)
		log.Debug(
			"Creating new object with id %s in %s",
			model.GetId(),
			model.Collection().FullName)

		err = model.Collection().Insert(model)
		if err != nil {
			log.Error(
				"Error creating model %s (%s): %s",
				model.GetId().Hex(),
				model.Collection().FullName, err)
		}
	} else {
		log.Debug(
			"Updating object with id %s in %s",
			model.GetId().Hex(),
			model.Collection().FullName)

		err = model.Collection().Update(bson.M{"_id": model.GetId()}, model)
		if err != nil {
			log.Error(
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

// Upsert the model
func Upsert(selector interface{}, model Modeller) (err error) {
	_, err = model.Collection().Upsert(selector, model)
	return
}
