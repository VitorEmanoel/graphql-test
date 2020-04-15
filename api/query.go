package api

import (
	"github.com/graphql-go/graphql"
	"graphql-test/api/product"
)

func NewQuery() *graphql.Object{
	var queryRoot = graphql.NewObject(graphql.ObjectConfig{Name:"query", Description: "Root query object", Fields: graphql.Fields{}})
	{
		product.NewProductQuery(queryRoot)
	}
	return queryRoot
}