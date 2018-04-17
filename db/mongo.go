package db

import (
	"gopkg.in/mgo.v2"
	"github.com/go-github/models"
	"time"
	"log"
)

type DB struct {
	Session *mgo.Session
}

func mgoDialInfo(mongo models.Mongo) *mgo.DialInfo {
	return &mgo.DialInfo{
		Addrs:          mongo.Addrs,
		Timeout:        10 * time.Second,
		Database:       mongo.Db,
	}
}

func (db *DB) Init(mongo models.Mongo)  {
	session, err := mgo.DialWithInfo(mgoDialInfo(mongo))
	if err != nil {
		log.Fatal("Couldn't connect to DB")
	}
	db.Session = session
}

func (db *DB) Close() {
	db.Close()
}