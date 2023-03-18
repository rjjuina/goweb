package main

import (
	"goweb/core/web/resolver"
	"goweb/core/web/schema"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	rootSchema := schema.MustGetRootSchema()
	var schemaOpts []graphql.SchemaOpt
	parseSchema := graphql.MustParseSchema(rootSchema, &resolver.Resolver{}, schemaOpts...)
	http.Handle("/query", &relay.Handler{Schema: parseSchema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
