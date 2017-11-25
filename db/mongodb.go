package db

import (
	"gopkg.in/mgo.v2"
	"github.com/iorellana18/Team-Udes-Back/utils"
)

var (
	mongoHost     string
	mongoDatabase string
)

func MongoSetup() {
	mongoHost = utils.Config.MongoDB.Host
	mongoDatabase = utils.Config.MongoDB.Database
}

func MongoSession() *mgo.Session {
	session, err := mgo.Dial(mongoHost)
	utils.Check(err)
	return session
}

func MongoDatabase(session *mgo.Session) *mgo.Database {
	return session.DB(mongoDatabase)
}

func MongoCollection(collection string, db *mgo.Database) *mgo.Collection {
	return db.C(collection)
}
