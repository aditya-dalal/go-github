package routes

import (
	"github.com/go-github/lib"
	"net/http"
	"github.com/go-github/models"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"fmt"
)

var actors = []route{
	{
		Name: "updateAvatar",
		Method: "PUT",
		Pattern: "",
		Handler: updateAvatarUrl,
	},
	{
		Name: "getActorsGroupByTotalEventsOrderByDesc",
		Method: "GET",
		Pattern: "",
		Handler: getActorsGroupByTotalEventsOrderByDesc,
	},
}

func updateAvatarUrl(c lib.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	var actor models.Actor
	var dbEvent models.DBEvent
	json.NewDecoder(r.Body).Decode(&actor)
	session := c.DB.Session.Copy()
	defer session.Copy()
	events := session.DB(c.Config.Mongo.Db).C(models.EventsCollection)
	err := events.Find(bson.M{"actor.id": actor.Id}).One(&dbEvent)
	if err != nil {
		return http.StatusNotFound, errors.New("actor not found")
	}
	if dbEvent.Actor.Login != actor.Login {
		fmt.Println(dbEvent.Actor.Login + "," + actor.Login)
		return http.StatusBadRequest, errors.New("cannot change login")
	}
	_, err = events.UpdateAll(bson.M{"actor.id": actor.Id}, bson.M{"$set": bson.M{"actor.avatar_url": actor.Avatar}})
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return http.StatusOK, nil
}

func getActorsGroupByTotalEventsOrderByDesc(c lib.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	var actors []models.Actor
	session := c.DB.Session.Copy()
	defer session.Copy()
	events := session.DB(c.Config.Mongo.Db).C(models.EventsCollection)
	pipe := []bson.M{
		{"$group": bson.M{"_id": "$actor.id", "count": bson.M{"$sum": 1}, "created_at": bson.M{"$max": "$created_at"}, "login": bson.M{"$first": "$actor.login"}, "avatar_url": bson.M{"$first": "$actor.avatar_url"}}},
		{"$sort": bson.M{"count": -1, "created_at": -1, "login": 1}},
		{"$project": bson.M{"_id": 0, "id": "$_id", "login": 1, "avatar_url":1}},
	}
	events.Pipe(pipe).All(&actors)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actors)
	return http.StatusOK, nil
}