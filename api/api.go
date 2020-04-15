package api

import (
	"github.com/graphql-go/graphql"
	handler "github.com/graphql-go/graphql-go-handler"
	"net/http"
)

func NewGraphQL() {
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: NewQuery(),
		Mutation: NewMutation(),
	})
	var graphqlHandler = handler.New(&handler.Config{
		Schema:           &schema,
		Pretty:           true,
		GraphiQL:         true,
		Playground:       true,
	})
	http.Handle("/graphql", graphqlHandler)
}
