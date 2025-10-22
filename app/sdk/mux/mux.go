// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"github.com/ardanlabs/service/app/domain/testapp"
	"github.com/ardanlabs/service/app/sdk/auth"
	"github.com/ardanlabs/service/app/sdk/mid"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

func WebAPI(log *logger.Logger, auth *auth.Auth) *web.App {
	app := web.NewApp(log.Info, mid.Logger(log), mid.Errors(log), mid.Panics())

	testapp.Routes(app, auth)

	return app
}
