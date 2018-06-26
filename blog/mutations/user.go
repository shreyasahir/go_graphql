package mutations

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"go_graphql/blog/db"
	"go_graphql/blog/types"
)

// GetCreateUserMutation creates a new user and returns it.
func GetCreateUserMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"firstname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			user := &types.User{
				Firstname: params.Args["firstname"].(string),
				Lastname:  params.Args["lastname"].(string),
			}

			// Add your user in database here
			conn := db.Connect()
			defer conn.Close()
			id := 0
			sqls := `INSERT INTO users (first_name,last_name) VALUES($1,$2) RETURNING id`
			err := conn.QueryRow(sqls, user.Firstname, user.Lastname).Scan(&id)
			if err != nil {
				panic(err)
			}
			user.ID = id
			return user, nil
		},
	}
}
