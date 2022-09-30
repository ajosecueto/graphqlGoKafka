package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"graphqlGoKafka/graph"
	"graphqlGoKafka/graph/generated"
	"graphqlGoKafka/utils"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {

	ctx := context.Background()
	go utils.StartConsumer(ctx)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
