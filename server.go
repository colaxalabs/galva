package main

import (
	"log"
	"net/http"

	"github.com/3dw1nM0535/galva/graph"
	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/utils"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

var port string

func init() {
	godotenv.Load()
	port = utils.MustGetEnv("PORT")
}

func main() {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.New()}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
