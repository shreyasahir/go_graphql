package types

import (
	"github.com/graphql-go/graphql"
)

//Owner type definition
type Owner struct {
	ID        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

//OwnerType is graphql schema for type owner
var OwnerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Owner",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"first_name": &graphql.Field{Type: graphql.String},
		"last_name":  &graphql.Field{Type: graphql.String},
	},
})
