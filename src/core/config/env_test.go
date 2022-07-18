package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentEnvironment(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(PROD_ENV, getCurrentEnvironment())
}

func TestGetCurrentEnvWithTestFlag(t *testing.T) {
	os.Args = append(os.Args, "--env=test")

	assert := assert.New(t)

	assert.Equal(TEST_ENV, getCurrentEnvironment())
}

func TestGetCurrentEnvWithDevFlag(t *testing.T) {
	os.Args = append(os.Args, "--env=dev")

	assert := assert.New(t)

	assert.Equal(DEV_ENV, getCurrentEnvironment())
}

func TestGetCurrentEnvWithStagingFlag(t *testing.T) {
	os.Args = append(os.Args, "--env=staging")

	assert := assert.New(t)

	assert.Equal(STAGING_ENV, getCurrentEnvironment())
}

func TestGetCurrentEnvWithInvalidFlag(t *testing.T) {
	os.Args = append(os.Args, "--env=stage")

	assert := assert.New(t)

	assert.Panics(func() { getCurrentEnvironment() })
}
