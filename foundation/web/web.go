// Package web asdasdas asd asd asd
package web

import (
	"context"
	"net/http"
)

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*http.ServeMux
	mw []MidFunc
}

func NewApp(mw ...MidFunc) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}
}

// HandleFunc IS MY API.
func (a *App) HandleFunc(pattern string, handlerFunc HandlerFunc, mw ...MidFunc) {
	handlerFunc = wrapMiddleware(mw, handlerFunc)
	handlerFunc = wrapMiddleware(a.mw, handlerFunc)

	h := func(w http.ResponseWriter, r *http.Request) {

		// DO WHAT I WANT

		handlerFunc(r.Context(), w, r)

		// DO WHAT I WANT
	}

	a.ServeMux.HandleFunc(pattern, h)
}
