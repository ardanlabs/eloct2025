package web

import (
	"context"

	uuid "github.com/google/uuid"
)

type ctxKey int

const (
	traceKey ctxKey = iota + 1
)

func setTraceID(ctx context.Context, traceID uuid.UUID) context.Context {
	return context.WithValue(ctx, traceKey, traceID)
}

func GetTraceID(ctx context.Context) uuid.UUID {
	v, ok := ctx.Value(traceKey).(uuid.UUID)
	if !ok {
		return uuid.UUID{}
	}

	return v
}
