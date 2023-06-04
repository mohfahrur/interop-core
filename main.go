package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/mohfahrur/interopCore/entity"
	usecase "github.com/mohfahrur/interopCore/usecase"
)

const schemaString = `
	schema {
		query: Query
	}
	type User {
		userID: ID!
		username: String!
		nim: String!
		githubLink:  String!
	}
	type Query {
		# List users:
		users: [User!]!
		# Get user:
		user(userID: ID!): User!
	}
`

var (
	opts   = []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	Schema = graphql.MustParseSchema(schemaString, &usecase.RootResolver{}, opts...)
)

func main() {
	ctx := context.Background()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from core",
		})
	})

	r.POST("/rent", func(c *gin.Context) {
		bodyByte, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"message": "error",
			})
			return
		}
		var req entity.FERequest
		err = json.Unmarshal(bodyByte, &req)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"message": "error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user":  req.Name,
			"email": req.Email,
			"hp":    req.HP,
		})
	})

	r.GET("/graphql", func(c *gin.Context) {
		var params entity.ClientQuery
		err := json.NewDecoder(c.Request.Body).Decode(&params)
		if err != nil {
			panic(err)
		}
		resp1 := Schema.Exec(ctx, params.Query, params.OpName, params.Variables)

		c.JSON(http.StatusOK, resp1)
	})
	r.Run()
}
