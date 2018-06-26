package app

import (
	"github.com/graphql-go/graphql"
)

func createUser(params graphql.ResolveParams) (interface{}, error) {
	ctx := params.Context
	name, _ := params.Args["name"].(string)
	user := &User{Name: name}
	key := datastore.NewIncompleteKey(ctx, "User", nil)
}
