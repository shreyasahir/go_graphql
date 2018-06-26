package queries

import (
	"github.com/graphql-go/graphql"
)

//GetRootFields returns all queries supported
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"owner":     GetOwnerQuery(),
		"lastowner": GetLastOwner(),
	}
}
