package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
	interopaD "github.com/mohfahrur/interop-core/domain/interopa"
	interopbD "github.com/mohfahrur/interop-core/domain/interopb"
	interopcD "github.com/mohfahrur/interop-core/domain/interopc"
	"github.com/mohfahrur/interop-core/entity"
	ticketUC "github.com/mohfahrur/interop-core/usecase/ticket"
	userUC "github.com/mohfahrur/interop-core/usecase/user"
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
	Schema = graphql.MustParseSchema(schemaString, &userUC.RootResolver{}, opts...)
)

func main() {
	ctx := context.Background()

	interopaDomain := interopaD.NewinteropaDomain()
	interopbDomain := interopbD.NewinteropbDomain()
	interopcDomain := interopcD.NewinteropcDomain()

	TicketUsercase := ticketUC.NewTicketUsecase(*interopaDomain,
		*interopbDomain,
		*interopcDomain)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from core",
		})
	})

	r.POST("/buy", func(c *gin.Context) {
		bodyByte, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"message": "error",
			})
			return
		}
		var req entity.FERequestGQL
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
			"item":  req.Item,
			"email": req.Email,
			"hp":    req.HP,
		})
	})

	r.POST("/v2/buy", func(c *gin.Context) {
		bodyByte, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"message": "error",
			})
			return
		}
		var req entity.User
		err = json.Unmarshal(bodyByte, &req)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"message": "error",
			})
			return
		}
		err = TicketUsercase.SendNotifikasi(req)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"message": "error",
			})
			return
		}
		err = TicketUsercase.UpdateData(req)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"message": "error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
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
	r.Run(":8000")
}
