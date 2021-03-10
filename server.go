package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/3dw1nM0535/galva/graph"
	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/store"
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

	orm, err := store.NewORM()
	if err != nil {
		fmt.Printf("Error while setting up ORM: " + err.Error())
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.New(orm)}))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		// send panic to sentry
		log.Print(err)
		debug.PrintStack()
		return errors.New("user message on panic")
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
