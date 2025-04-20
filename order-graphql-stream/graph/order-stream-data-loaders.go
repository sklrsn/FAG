package graph

import (
	"context"
	"time"

	"github.com/vikstrous/dataloadgen"
)

type DataLoaders struct {
	OrderLoader *dataloadgen.Loader[string, Record]
}

func NewDataLoaders(dbConn DbConn) *DataLoaders {
	return &DataLoaders{
		OrderLoader: dataloadgen.NewLoader(func(ctx context.Context, keys []string) ([]Record, []error) {
			id := ""
			if len(keys) > 0 {
				id = keys[0]
			}
			r, err := dbConn.Order(id)
			return []Record{r}, []error{err}
		}, dataloadgen.WithWait(200*time.Millisecond)),
	}
}
