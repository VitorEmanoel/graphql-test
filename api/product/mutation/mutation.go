package mutation

import (
	"github.com/graphql-go/graphql"
)


func NewMutation(mutationRoot *graphql.Object) {
	mutationRoot.AddFieldConfig("createProduct", &CreateProductField)
}
