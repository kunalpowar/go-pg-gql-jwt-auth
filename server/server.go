package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kunalpowar/gopggqlauth/server/auth"
	"github.com/kunalpowar/gopggqlauth/server/db"
	"github.com/kunalpowar/gopggqlauth/server/graph"
	"github.com/kunalpowar/gopggqlauth/server/graph/generated"

	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	ctx := context.Background()
	db.MustInit(ctx)

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
