package usecase

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/mohfahrur/interop-core/entity"
)

// Define mock data:
var users = []entity.User{
	{
		UserID:     graphql.ID("u-001"),
		Username:   "Rizky Wahyu Dewantoro",
		Nim:        "4.33.21.1.23",
		GithubLink: "https://github.com/Kyroline/intero-api-1",
	},
	{
		UserID:     graphql.ID("u-002"),
		Username:   "Fido Jahfal Prayoga",
		Nim:        "4.33.21.1.10",
		GithubLink: "https://github.com/fidojahfal/simpleApi-HelloWorld",
	},
	{
		UserID:     graphql.ID("u-003"),
		Username:   "Bhatinden Sejiarga E. G.",
		Nim:        "4.33.21.1.06",
		GithubLink: "https://github.com/nekoeyes26/interoperabilitas_api_1.git",
	},
	{
		UserID:     graphql.ID("u-004"),
		Username:   "Muhammad David Akbar",
		Nim:        "43321118",
		GithubLink: "https://github.com/davidRT2/simple-api.git",
	},
	{
		UserID:     graphql.ID("u-005"),
		Username:   "Ravi Naufal",
		Nim:        "43321121",
		GithubLink: "https://github.com/veenss/interoperabilitas_2",
	},
}

type RootResolver struct{}

func (r *RootResolver) Users() ([]entity.User, error) {
	return users, nil
}

func (r *RootResolver) User(args struct{ UserID graphql.ID }) (entity.User, error) {
	for _, user := range users {
		if args.UserID == user.UserID {
			return user, nil
		}
	}
	return entity.User{}, nil
}
