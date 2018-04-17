package routes

import (
	"net/http"
	"github.com/go-github/models"
	"encoding/json"
	"github.com/go-github/lib"
	"errors"
	"time"
	"gopkg.in/mgo.v2"
)

var events = []route{
	{
		Name: "allEvents",
		Method: "GET",
		Pattern: "/",
		Handler: getAllEvents,
	},

	{
		Name: "create",
		Method: "POST",
		Pattern: "/",
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
	var eventList = []models.Event{}
	events.Find(nil).Sort("_id").All(&eventList)
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
	return http.StatusCreated, nil
}

func getAllEventsForActor(c lib.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return http.StatusOK, nil
}