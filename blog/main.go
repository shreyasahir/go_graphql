package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"go_graphql/blog/mutations"
	"go_graphql/blog/queries"
	"log"
	"net/http"
)

func main() {
	log.Print("ready: listening...\n")
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
		log.Fatalf("Failed to create new schema, error: %v", err)
	}
	httpHandler := handler.New(&handler.Config{
		Schema: &schema,
	})
	fmt.Println("got request")

	http.Handle("/", httpHandler)
	http.ListenAndServe(":8000", nil)

}
