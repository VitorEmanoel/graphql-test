package mutation

import (
	"github.com/graphql-go/graphql"
	"graphql-test/api/product/types"
)

var CreateProductField = graphql.Field{
	Name:              "Create Product",
	Type:              types.ProductType,
	Args:			   graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type:         graphql.NewNonNull(graphql.String),
			Description:  "Name of product",
		},
		"info": &graphql.ArgumentConfig{
			Type:         graphql.String,
			Description:  "Info of product",
		},
		"price": &graphql.ArgumentConfig{
			Type:         graphql.NewNonNull(graphql.Float),
			Description:  "Price of product",
		},
	},
	Resolve:           ProductsResolver,
	Description:       "Create Product",
}