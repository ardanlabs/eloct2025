// Package web asdasdas asd asd asd
package web

import (
	"context"
	"net/http"
)

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*http.ServeMux
}

func NewApp() *App {
	return &App{
		ServeMux: http.NewServeMux(),
	}
}

// HandleFunc IS MY API.
func (app *App) HandleFunc(pattern string, handler HandlerFunc) {
	h := func(w http.ResponseWriter, r *http.Request) {

		// DO WHAT I WANT

		handler(r.Context(), w, r)

		// DO WHAT I WANT
	}

	app.ServeMux.HandleFunc(pattern, h)
}
