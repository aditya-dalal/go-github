package models

type Repo struct {
	Id   int    `bson:"id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
	Url  string `bson:"url" json:"url"`
}
