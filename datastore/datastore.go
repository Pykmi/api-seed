package datastore

import (
	"context"
	"net/http"

	gocb "gopkg.in/couchbase/gocb.v1"
)

type StoreOptions struct {
	DBPath        string
	Namespace     string
	RetryAttempts int
}

type Store struct {
	Cluster *gocb.Cluster
	Bucket  *gocb.Bucket
	Options StoreOptions
}

func Middleware(store *Store) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "STORE", store)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func New(options StoreOptions) (*Store, error) {
	store := &Store{}

	cluster, err := gocb.Connect(options.DBPath)
	if err != nil {
		return nil, err
	}

	bucket, err := cluster.OpenBucket(options.Namespace, "")
	if err != nil {
		return nil, err
	}

	store.Cluster = cluster
	store.Bucket = bucket
	store.Options = options
	return store, nil
}
