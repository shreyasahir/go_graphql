package mutations

import (
	"github.com/graphql-go/graphql"
	"go_graphql/petstore/db"
	"go_graphql/petstore/types"
	"log"
	"strconv"
)

//GetCreateOwnerMutation will create owner and return the created one
func GetCreateOwnerMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.OwnerType,
		Args: graphql.FieldConfigArgument{
			"first_name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"last_name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			owner := types.Owner{
				FirstName: params.Args["first_name"].(string),
				LastName:  params.Args["last_name"].(string),
			}
			conn := db.Connect()
			defer conn.Close()
			id := 0
			sqls := `INSERT INTO owner (first_name,last_name) VALUES($1,$2) RETURNING id`
			err := conn.QueryRow(sqls, owner.FirstName, owner.LastName).Scan(&id)
			if err != nil {
				panic(err)
			}
			owner.ID = id
			return owner, nil
		},
	}

}

//GetUpdateOwnerMutation updates owner and returns the object
func GetUpdateOwnerMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.OwnerType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"first_name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"last_name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := strconv.Atoi(params.Args["id"].(string))
			var firstname = ""
			if _, ok := params.Args["first_name"]; ok {
				firstname = params.Args["first_name"].(string)
			} else {
				firstname = ""
			}

			var sqls string
			var owner types.Owner
			var err error
			conn := db.Connect()
			defer conn.Close()
			log.Println("id is", id)
			if firstname != "" {
				owner = types.Owner{
					ID:        id,
					FirstName: firstname,
					LastName:  params.Args["last_name"].(string),
				}
				sqls = `UPDATE owner SET first_name = $1,last_name=$2 where id=$3`
				_, err = conn.Exec(sqls, owner.FirstName, owner.LastName, id)

			} else {
				owner = types.Owner{
					ID:       id,
					LastName: params.Args["last_name"].(string),
				}
				sqls = `UPDATE owner SET last_name=$1 where id=$2`
				_, err = conn.Exec(sqls, owner.LastName, id)

			}

			log.Println("queried db", err)
			if err != nil {
				panic(err)
			}

			return owner, nil
		},
	}
}
