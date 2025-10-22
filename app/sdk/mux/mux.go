// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"encoding/json"
	"net/http"
)

func WebAPI() *http.ServeMux {
	mux := http.NewServeMux()

	h := func(w http.ResponseWriter, r *http.Request) {
		v := struct {
			Status string
		}{
			Status: "OK",
		}

		json.NewEncoder(w).Encode(v)
	}

	mux.HandleFunc("/test", h)

	return mux
}
