package query

import (
	"github.com/graphql-go/graphql"
	"graphql-test/repository"
)

var ProductsResolver = func (p graphql.ResolveParams) (interface{}, error) {
	return repository.PublicProductRepository.FindAll(), nil
}

var ProductResolver = func(p graphql.ResolveParams) (interface{}, error) {
	return repository.PublicProductRepository.Find(p.Args["id"].(int)), nil
}
