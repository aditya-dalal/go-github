package models


type Config struct {
	Server Server
	Mongo Mongo
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Mongo struct {
	Addrs []string `yaml:"addrs"`
	Db string `yaml:"db"`
}
