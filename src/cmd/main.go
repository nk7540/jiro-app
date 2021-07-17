package main

import (
	"flag"
	"log"
	"net/http"

	"artics-api/src/config"
	"artics-api/src/internal/graph"
	"artics-api/src/internal/graph/generated"
	"artics-api/src/internal/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

const defaultPort = "8080"

func main() {
	// Initialization
	configFile := flag.String("config", "config.yml", "User Config file")
	flag.Parse()
	app := &config.AppConfig{ConfigFile: *configFile}
	app.Setup()
	log.Print("Setup complete")

	// Running application
	router := chi.NewRouter()

	router.Use(middleware.Auth(&app.Database, &app.Auth))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(
		&app.Uploader,
		&app.Auth,
		&app.Mail,
		&app.Database,
		&app.RPC,
	)}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":8080", router))
}
