package main

import (
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
	opts := db.InitDBOpts{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
	}

	tcpHost := os.Getenv("DB_TCP_HOST")
	if tcpHost != "" {
		opts.Network = "tcp"
		opts.Addr = tcpHost
	} else {
		dbSocketName := os.Getenv("INSTANCE_CONNECTION_NAME")
		if dbSocketName == "" {
			log.Fatalf("at least one of DB_TCP_HOST or INSTANCE_CONNECTION_NAME should be set in env")
		}

		opts.Network = "unix"
		opts.Addr = dbSocketName
	}

	db.MustInit(&opts)

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
