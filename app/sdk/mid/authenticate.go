package mid

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/app/sdk/auth"
	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
	"github.com/google/uuid"
)

func Bearer(ath *auth.Auth) web.MidFunc {
	m := func(next web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, r *http.Request) web.Encoder {
			claims, err := ath.Authenticate(ctx, r.Header.Get("authorization"))
			if err != nil {
				return errs.New(errs.Unauthenticated, err)
			}

			if claims.Subject == "" {
				return errs.Newf(errs.Unauthenticated, "authorize: you are not authorized for that action, no claims")
			}

			subjectID, err := uuid.Parse(claims.Subject)
			if err != nil {
				return errs.Newf(errs.Unauthenticated, "parsing subject: %s", err)
			}

			ctx = setUserID(ctx, subjectID)
			ctx = setClaims(ctx, claims)

			return next(ctx, r)
		}

		return h
	}

	return m
}
