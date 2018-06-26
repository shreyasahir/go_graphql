package queries

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"go_graphql/blog/db"
	"go_graphql/blog/types"
)

// GetUserQuery returns the queries available against user type.
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var users []types.User
			user := new(types.User)
			conn := db.Connect()
			rows, err := conn.Query("SELECT * FROM users")
			if err != nil {
				fmt.Println("Problem quering db")
			}
			for rows.Next() {
				err = rows.Scan(&user.ID, &user.Firstname, &user.Lastname)
				if err != nil {
					fmt.Println("Problem quering db")
				}
				users = append(users, *user)
			}
			// ... Implement the way you want to obtain your data here.
			defer conn.Close()

			return users, nil
		},
	}
}
