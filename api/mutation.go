package api

import (
	"github.com/graphql-go/graphql"
	"graphql-test/api/product"
)

func NewMutation() *graphql.Object{
	var mutationRoot = graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Description: "Root Mutation", Fields: graphql.Fields{}})
	{
		product.NewProductMutation(mutationRoot)
	}
	return mutationRoot
}
