package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDbConnectContext(t *testing.T) {
	assert := assert.New(t)

	dbContext := GetDbConnectContext()

	assert.Equal(DatabaseConnection, dbContext.Purpose)
	assert.NotNil(dbContext.Ctx)
}

func TestGetDbRWContext(t *testing.T) {
	assert := assert.New(t)

	dbContext := GetDbRWContext()

	assert.Equal(DatabaseReadWrite, dbContext.Purpose)
	assert.NotNil(dbContext.Ctx)
}
