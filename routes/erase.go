package routes

import (
	"net/http"
	"github.com/go-github/lib"
	"github.com/go-github/models"
)

var erase = []route{
	{
		Name: "erase",
		Method: "DELETE",
		Pattern: "/",
		Handler: eraseAllEvents,
	},
}

func eraseAllEvents(c lib.AppContext, w http.ResponseWriter, r *http.Request) (int, error){
	session := c.DB.Session.Copy()
	defer session.Close()
	events := session.DB(c.Config.Mongo.Db).C(models.EventsCollection)
	events.Remove(nil)
	return http.StatusOK, nil
}