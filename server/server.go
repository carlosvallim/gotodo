package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/carlosvallim/gotodo"
	"github.com/carlosvallim/gotodo/data"
)

func init() {
	data.Db = data.ConnectPostgres()
	fmt.Println("Connection with database has a successful!")
}

//const defaultPort = "8080"

func main() {
	/*port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}*/

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gotodo.NewExecutableSchema(gotodo.Config{Resolvers: &gotodo.Resolver{}})))

	log.Println("connect for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
