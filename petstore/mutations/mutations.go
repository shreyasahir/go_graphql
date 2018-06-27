package mutations

import (
	"github.com/graphql-go/graphql"
)

//GetRootFields returns all queries supported
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"createOwner": GetCreateOwnerMutation(),
		"updateOwner": GetUpdateOwnerMutation(),
	}
}
