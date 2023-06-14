package entity

import "github.com/graph-gophers/graphql-go"

type UserGQL struct {
	UserID     graphql.ID
	Username   string `json:"user"`
	Nim        string `json:"email"`
	GithubLink string `json:"item"`
}

type User struct {
	User  string `json:"user"`
	Email string `json:"email"`
	Item  string `json:"item"`
	Hp    string `json:"hp"`
}
