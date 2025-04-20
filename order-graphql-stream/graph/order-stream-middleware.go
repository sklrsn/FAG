package graph

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
)

type MiddlewareKey string

const KeyOrderMiddleware MiddlewareKey = MiddlewareKey("key-order-middleware")

func For(ctx context.Context) *DataLoaders {
	return ctx.Value(KeyOrderMiddleware).(*DataLoaders)
}
func OrderMiddleware(dbConn DbConn, next *handler.Server) http.Handler {
	loaders := NewDataLoaders(dbConn)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), KeyOrderMiddleware, loaders))
		next.ServeHTTP(w, r)
	})
}
