package models

import "time"

const EventsCollection = "events"

type Event struct {
	Id        int       `bson:"_id,omitempty" json:"id"`
	Type      string    `bson:"type" json:"type"`
	Actor     Actor     `bson:"actor" json:"actor"`
	Repo      Repo      `bson:"repo" json:"repo"`
	CreatedAt JsonTime `bson:"created_at" json:"created_at,time"`
}

type DBEvent struct {
	Id        int       `bson:"_id,omitempty" json:"id"`
	Type      string    `bson:"type" json:"type"`
	Actor     Actor     `bson:"actor" json:"actor"`
	Repo      Repo      `bson:"repo" json:"repo"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}