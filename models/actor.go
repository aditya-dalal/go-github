package models

type Actor struct {
	Id     int    `bson:"id,omitempty" json:"id"`
	Login  string `bson:"login" json:"login"`
	Avatar string `bson:"avatar_url" json:"avatar_url"`
}
