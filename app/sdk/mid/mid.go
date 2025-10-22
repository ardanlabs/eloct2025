// Package mid asdas das dsa dsq.
package mid

import (
	"context"
	"errors"

	"github.com/ardanlabs/service/app/sdk/auth"
	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
	"github.com/google/uuid"
)

type ctxKey int

const (
	claimKey ctxKey = iota + 1
	userIDKey
)

func isError(e web.Encoder) error {
	err, isError := e.(error)
	if isError {
		var appErr *errs.Error
		if errors.As(err, &appErr) && appErr.Code == errs.None {
			return nil
		}
		return err
	}
	return nil
}

func setClaims(ctx context.Context, claims auth.Claims) context.Context {
	return context.WithValue(ctx, claimKey, claims)
}

// GetClaims returns the claims from the context.
func GetClaims(ctx context.Context) auth.Claims {
	v, ok := ctx.Value(claimKey).(auth.Claims)
	if !ok {
		return auth.Claims{}
	}
	return v
}

func setUserID(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// GetUserID returns the user id from the context.
func GetUserID(ctx context.Context) (uuid.UUID, error) {
	v, ok := ctx.Value(userIDKey).(uuid.UUID)
	if !ok {
		return uuid.UUID{}, errors.New("user id not found in context")
	}

	return v, nil
}
