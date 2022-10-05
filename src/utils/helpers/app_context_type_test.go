package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppContextCancel(t *testing.T) {
	assert := assert.New(t)

	appCtx := GetDbConnectContext()

	assert.NotPanics(func() { appCtx.Cancel() })
}
