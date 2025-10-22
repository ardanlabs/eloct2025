// Package web asdasdas asd asd asd
package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type Logger func(ctx context.Context, msg string, args ...any)

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

type HandlerFunc func(ctx context.Context, r *http.Request) Encoder

type App struct {
	*http.ServeMux
	log Logger
	mw  []MidFunc
}

func NewApp(log Logger, mw ...MidFunc) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		log:      log,
		mw:       mw,
	}
}

// HandleFunc IS MY API.
func (a *App) HandleFunc(pattern string, handlerFunc HandlerFunc, mw ...MidFunc) {
	handlerFunc = wrapMiddleware(mw, handlerFunc)
	handlerFunc = wrapMiddleware(a.mw, handlerFunc)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := setTraceID(r.Context(), uuid.New())

		// DO WHAT I WANT

		resp := handlerFunc(ctx, r)

		if err := Respond(ctx, w, resp); err != nil {
			a.log(ctx, "web-respond", "ERROR", err)
			return
		}

		// DO WHAT I WANT
	}

	a.ServeMux.HandleFunc(pattern, h)
}
