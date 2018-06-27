package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"go_graphql/petstore/mutations"
	"go_graphql/petstore/queries"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.OpenFile("development.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Unable to open file for logging", err)
	}
	defer f.Close()
	log.SetOutput(f)
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutations.GetRootFields(),
		}),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("Failed to create new schema ", err, err.Error())
	}
	httpHandler := handler.New(&handler.Config{
		Schema: &schema,
	})

	http.Handle("/", httpHandler)
	http.ListenAndServe(":8080", nil)
}
