package entity

import "github.com/graph-gophers/graphql-go"

type User struct {
	UserID     graphql.ID
	Username   string
	Nim        string
	GithubLink string
}
