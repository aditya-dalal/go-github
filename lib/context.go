package lib

import (
	"github.com/go-github/db"
	"github.com/go-github/models"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type AppContext struct {
	DB db.DB
	Config models.Config
}

func (c *AppContext) InitDB(mongoConfig models.Mongo) {
	c.DB.Init(mongoConfig)
}

func (c *AppContext) CloseDB() {
	c.DB.Close()
}

func (c *AppContext) LoadConfig(configFile string)  {
	var yamlFile, err = ioutil.ReadFile(configFile)
	if err != nil {
	log.Fatalf("Error reading config file %q: %s\n", configFile, err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &c.Config)
	if err != nil {
	log.Fatalf("Error unmarshalling config file %q: %s\n", configFile, err.Error())
	}
}