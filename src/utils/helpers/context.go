package helpers

import (
	"context"
	"time"
)

const (
	// Timeout operations after N seconds
	connectTimeout   = 5
	readWriteTimeout = 3
)

func GetDbConnectContext() *AppContext {
	ctx := &AppContext{Purpose: DatabaseConnection}
	return _getContext(ctx, connectTimeout)
}

func GetDbRWContext() *AppContext {
	ctx := &AppContext{Purpose: DatabaseReadWrite}
	return _getContext(ctx, readWriteTimeout)
}

// Will return the context
func _getContext(ctx *AppContext, timeout time.Duration) *AppContext {
	ctx.Ctx, ctx.cancelFunc = context.WithTimeout(context.Background(), timeout*time.Second)
	if ctx.Ctx == nil {
		panic("Failed to get context")
	}
	return ctx
}
