package context

import (
	"context"
	"fmt"
	"net/http"
)

// 大量にプロセスを消費する。
// 一貫した停止方法が大事

// type Store interface {
// 	Fetch() string
// 	Cancel()
// }

// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		  store.Cancel()
// 			fmt.Fprint(w, store.Fetch())
// 	}
// }

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			data, err := store.Fetch(r.Context())

			if err != nil {
					return // todo: log error however you like
			}

			fmt.Fprint(w, data)
	}
}