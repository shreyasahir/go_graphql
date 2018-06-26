package queries

import (
	"github.com/graphql-go/graphql"
	"go_graphql/petstore/db"
	"go_graphql/petstore/types"
	"log"
	"strconv"
)

//GetOwnerQuery queries and replies with single query
func GetOwnerQuery() *graphql.Field {
	return &graphql.Field{
		Type:        types.OwnerType,
		Description: "Get single Owner",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			idQuery, _ := params.Args["id"].(string)
			log.Println("Got request for Owner with id ", idQuery)
			owner := new(types.Owner)
			conn := db.Connect()
			defer db.Connect()
			id, _ := strconv.Atoi(idQuery)
			rows, err := conn.Query("SELECT id,first_name,last_name from OWNER WHERE id=$1", id)
			if err != nil {
				log.Fatal("Encountered error while fetching data", err)
			}
			for rows.Next() {
				err := rows.Scan(&owner.ID, &owner.FirstName, &owner.LastName)
				if err != nil {
					log.Fatal("Encountered error while looping data", err)
				}
			}
			return owner, nil
		},
	}
}

//GetLastOwner returns last owner in our registry
func GetLastOwner() *graphql.Field {
	return &graphql.Field{
		Type:        types.OwnerType,
		Description: "Get Single Last owner registered",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			conn := db.Connect()
			defer conn.Close()
			owner := new(types.Owner)
			rows, err := conn.Query("SELECT id,first_name,last_name from OWNER order by id DESC limit 1")
			if err != nil {
				log.Fatalln("Error while retrieving last owner", err)
			}
			for rows.Next() {
				err := rows.Scan(&owner.ID, &owner.FirstName, &owner.LastName)
				if err != nil {
					log.Fatal("Encountered error while looping data", err)
				}
			}
			log.Printf("%v", owner)
			return owner, nil
		},
	}
}
