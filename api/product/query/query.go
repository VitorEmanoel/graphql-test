package query

import "github.com/graphql-go/graphql"

func NewQuery(queryRoot *graphql.Object) {
	queryRoot.AddFieldConfig("products", &ProductsField)
	queryRoot.AddFieldConfig("product", &ProductField)
}