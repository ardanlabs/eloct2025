// Package testapp is for class.
package testapp

import (
	"context"
	"encoding/json"
	"net/http"
)

func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v := struct {
		Status string
	}{
		Status: "OK",
	}

	return json.NewEncoder(w).Encode(v)
}
