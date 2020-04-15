package repository

import "graphql-test/models"

type ProductRepository interface {
	Find(id int) models.Product
	FindAll() []models.Product
	Create(product models.Product) models.Product
}

type ProductRepositoryContext struct {
	Products	[]models.Product
}

func (p *ProductRepositoryContext) Find(id int) models.Product {
	for _, product := range p.Products {
		if product.Id == id {
			return product
		}
	}
	return models.Product{}
}

func (p *ProductRepositoryContext) FindAll() []models.Product {
	return p.Products
}

func (p *ProductRepositoryContext) Create(product models.Product) models.Product {
	product.Id = len(p.Products) + 1
	p.Products = append(p.Products, product)
	return product
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryContext{}
}
