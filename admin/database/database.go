package database

import (
	"github.com/sebdah/recharged/central-system/settings"
	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

// Connect to MongoDB
func GetSession() *mgo.Session {
	if Session == nil {
		conf := settings.GetSettings()
		session, err := mgo.Dial(conf.MongoDBHosts)
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

	conf := settings.GetSettings()

	return Session.DB(conf.DatabaseName)
}
