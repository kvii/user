package main

import (
	"fmt"
	"net/http"

	"github.com/kvii/pkg/database"
	"github.com/kvii/user/internal/user"
)

func main() {
	s := user.Service{
		DB: &database.DB{
			Source: "user",
		},
	}
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		name := r.URL.Query().Get("name")

		msg, err := s.Hello(ctx, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			fmt.Fprintln(w, msg)
		}
	})
	http.ListenAndServe(":9090", nil)
}
