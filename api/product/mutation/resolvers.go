package mutation

import (
	"github.com/graphql-go/graphql"
	"graphql-test/models"
	"graphql-test/repository"
)

var ProductsResolver = func (p graphql.ResolveParams) (interface{}, error) {
	var product = models.Product{
		Name:  p.Args["name"].(string),
		Price: p.Args["price"].(float64),
	}
	if info := p.Args["info"]; info != nil{
		product.Info = info.(string)
	}
	resultProduct := repository.PublicProductRepository.Create(product)
	return resultProduct, nil
}
