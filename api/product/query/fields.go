package query

import (
	"github.com/graphql-go/graphql"
	"graphql-test/api/product/types"
)

var ProductsField = graphql.Field{
	Name:              "Products",
	Type:              graphql.NewList(types.ProductType),
	Resolve:           ProductsResolver,
	Description:       "Find all products",
}

var ProductField = graphql.Field{
	Name:              "Product",
	Type:              types.ProductType,
	Args:              graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:         graphql.NewNonNull(graphql.Int),
			Description:  "ID for search product",
		},
	},
	Resolve:           ProductResolver,
	Description:       "Find specify product by id",
}