package routes

import (
	"net/http"
	"github.com/go-github/models"
	"encoding/json"
	"github.com/go-github/lib"
	"errors"
	"time"
	"gopkg.in/mgo.v2"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

var events = []route{
	{
		Name: "allEvents",
		Method: "GET",
		Pattern: "",
		Handler: getAllEvents,
	},

	{
		Name: "create",
		Method: "POST",
		Pattern: "",
		Handler: createEvent,
	},

	{
		Name: "allEventsForActor",
		Method: "GET",
		Pattern: "/actors/{actorID}",
		Handler: getAllEventsForActor,
	},
}

func getAllEvents(c lib.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	session := c.DB.Session.Copy()
	defer session.Close()
	events := session.DB(c.Config.Mongo.Db).C(models.EventsCollection)
	var dbEvents = []models.DBEvent{}
	var eventList = []models.Event{}
	events.Find(nil).Sort("_id").All(&dbEvents)
	for _, event := range dbEvents {
		eventList = append(eventList, models.Event{
			Id: event.Id,
			Type: event.Type,
			Actor: event.Actor,
			Repo: event.Repo,
			CreatedAt: models.JsonTime(event.CreatedAt.UTC()),
		})
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(eventList)
	return http.StatusOK, nil
}

func createEvent(c lib.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	var event models.Event
	json.NewDecoder(r.Body).Decode(&event)
	session := c.DB.Session.Copy()
	defer session.Close()
	events := session.DB(c.Config.Mongo.Db).C(models.EventsCollection)
	err := events.Insert(models.DBEvent{
		Id: event.Id,
		Type: event.Type,
		Actor: event.Actor,
		Repo: event.Repo,
		CreatedAt: time.Time(event.CreatedAt),
	})
	if err != nil {
		if mgo.IsDup(err) {
			return http.StatusBadRequest, errors.New("Event already exists")
		}
		return http.StatusInternalServerError, errors.New("Failed to insert event to db")
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return http.StatusCreated, nil
}

func getAllEventsForActor(c lib.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	session := c.DB.Session.Copy()
	defer session.Close()
	events := session.DB(c.Config.Mongo.Db).C(models.EventsCollection)
	vars := mux.Vars(r)
	actorId, _ := strconv.Atoi(vars["actorID"])
	var eventList = []models.Event{}
	err := events.Find(bson.M{
		"actor.id": actorId,
	}).Sort("_id").All(&eventList)
	if err != nil {
		return http.StatusInternalServerError, errors.New("Failed to insert event to db")
	}
	if len(eventList) == 0 {
		return http.StatusNotFound, errors.New("Actor not found")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(eventList)
	return http.StatusOK, nil
}