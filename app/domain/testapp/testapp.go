// Package testapp is for class.
package testapp

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

func Test(ctx context.Context, r *http.Request) web.Encoder {
	v := status{
		Status: "OK 2",
	}

	return v
}
