package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sklrsn/FAG/orders-graphql-stream/graph"
	"github.com/vektah/gqlparser/v2/ast"
)

const streamPort = "9094"

func main() {
	dbConn := new(graph.DbStore).Connect()
	srv := handler.New(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.OrdersResolver{
					DbConn: dbConn,
				},
			}))

	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	srv.Use(extension.FixedComplexityLimit(20))

	router := mux.NewRouter()
	router.Handle("/graphql", graph.OrderMiddleware(dbConn, srv))

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", streamPort),
		Handler: router,
	}

	log.Println("starting order graphql engine")
	log.Fatalf("%v", server.ListenAndServe())
}
