package testapp

import (
	"github.com/ardanlabs/service/app/sdk/auth"
	"github.com/ardanlabs/service/app/sdk/mid"
	"github.com/ardanlabs/service/foundation/web"
)

func Routes(app *web.App, auth *auth.Auth) {
	authMid := mid.Bearer(auth)

	app.HandleFunc("/test", test)
	app.HandleFunc("/testauth", test, authMid)
}
