package product

import (
	"github.com/graphql-go/graphql"
	"graphql-test/api/product/mutation"
	"graphql-test/api/product/query"
)

func NewProductQuery(queryRoot *graphql.Object){
	query.NewQuery(queryRoot)
}

func NewProductMutation(mutationRoot *graphql.Object) {
	mutation.NewMutation(mutationRoot)
}