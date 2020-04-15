package repository

var PublicProductRepository ProductRepository

func init() {
	PublicProductRepository = NewProductRepository()
}
