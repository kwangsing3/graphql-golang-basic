package main

import (
	"log"
	"net/http"

	"github.com/kwangsing3/graphql-golang-basic/dbhandler"
	"github.com/kwangsing3/graphql-golang-basic/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "80"

func main() {
	// mongodb
	defer dbhandler.DisConnect()

	//web service
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
