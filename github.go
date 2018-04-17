package main

import (
	"github.com/go-github/routes"
	"net/http"
	"path/filepath"
	"os"
	"github.com/go-github/lib"
)

func main() {
	absConfigPath, _ := filepath.Abs(os.Args[1])
	var app = &lib.AppContext{}
	app.LoadConfig(absConfigPath)

	app.InitDB(app.Config.Mongo)
	defer app.CloseDB()

	router := routes.CreateRouter(*app)
	http.ListenAndServe(":" + app.Config.Server.Port, router)
}
