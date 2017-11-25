package db

import (
	"gopkg.in/mgo.v2"

	"github.com/citiaps/aysana-backend/utils"
)

var (
	mongoHost     string
	mongoDatabase string
	mongoUsername string
	mongoPass string
)

func MongoSetup() {
	mongoHost = utils.Config.MongoDB.Host
	mongoDatabase = utils.Config.MongoDB.Db
	mongoUsername = utils.Config.MongoDB.Username
	mongoPass = utils.Config.MongoDB.Pass
}

func MongoSession() *mgo.Session {
	info := &mgo.DialInfo{
		Addrs:    []string{mongoHost},

		Database: mongoDatabase,
		Username: mongoUsername,
		Password: mongoPass,
	}

	session, err := mgo.DialWithInfo(info)
	utils.Check(err)
	return session
}

func MongoDatabase(session *mgo.Session) *mgo.Database {
	return session.DB(mongoDatabase)
}

func MongoCollection(collection string, db *mgo.Database) *mgo.Collection {
	return db.C(collection)
}
