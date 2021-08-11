package helpers

import (
	"context"
	"time"
)

const (
	// Timeout operations after N seconds
	connectTimeout = 5
)

// Will return the context
func GetContext() context.Context {
	// Get context
	ctx, _ := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	if ctx == nil {
		panic("Failed to get context")
	}
	return ctx
}
