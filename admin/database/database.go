package database

import (
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"github.com/sebdah/recharged/admin/config"
)

var (
	Session *mgo.Session
	log     logging.Logger
)

// Connect to MongoDB
func GetSession() *mgo.Session {
	if Session == nil {
		session, err := mgo.Dial(config.Config.GetString("mongodb.hosts"))
		if err != nil {
			panic(err)
		}
		Session = session
	}

	return Session
}

// Get the database
func GetDb() *mgo.Database {
	if Session == nil {
		GetSession()
	}

	return Session.DB(config.Config.GetString("mongodb.db"))
}

// Ensure all databases
func EnsureAllDatabases() {
	if config.Env == "dev" {
		log.Info("Ensuring databases")
		CreateCollectionBootNotificationLogs()
		CreateCollectionChargePoints()
		CreateCollectionIdTags()
	}
}
