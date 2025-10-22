// Package testapp is for class.
package testapp

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
)

func Test(ctx context.Context, r *http.Request) web.Encoder {
	if n := rand.Intn(100); n%2 == 0 {
		//panic("OHH NOOOO WHYYYYYYY!")
		return errs.Newf(errs.InvalidArgument, "Customer Error")
	}

	v := status{
		Status: "OK 2",
	}

	return v
}
