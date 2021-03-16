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
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var port string

func init() {
	godotenv.Load()
	port = utils.MustGetEnv("PORT")
}

func main() {

	router := chi.NewRouter()

	// Add CORS middleware on every request
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	orm, err := store.NewORM()
	if err != nil {
		fmt.Printf("Error while setting up ORM: " + err.Error())
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.New(orm)}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		// send panic to sentry
		log.Print(err)
		debug.PrintStack()
		return errors.New("user message on panic")
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
