package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	mockDbName     = "db_project"
	mockDbHost     = "mongo.db.localhost"
	mockDbPort     = "27017"
	mockDbUsername = "db_user1"
	mockDbPassword = "db_password*#12&^"
)

func _setDbConfigMocks(t *testing.T) {
	t.Setenv(dBName, mockDbName)
	t.Setenv(dBHost, mockDbHost)
	t.Setenv(dBPort, mockDbPort)
	t.Setenv(dBUsername, mockDbUsername)
	t.Setenv(dBPassword, mockDbPassword)
}

func TestLoadDbConfig(t *testing.T) {
	assert := assert.New(t)

	_setDbConfigMocks(t)

	result := loadDbConfig(PROD_ENV)

	assert.Equal(mockDbName, result.DbName)
	assert.Equal(mockDbHost, result.DbHost)
	assert.Equal(mockDbPort, result.DbPort)
	assert.Equal(mockDbUsername, result.DbUsername)
	assert.Equal(mockDbPassword, result.DbPassword)
}

func TestLoadDbConfigDbNames(t *testing.T) {
	assert := assert.New(t)

	_setDbConfigMocks(t)

	result := loadDbConfig(PROD_ENV)
	assert.Equal(mockDbName, result.DbName)

	result = loadDbConfig(DEV_ENV)
	assert.Equal(fmt.Sprintf(DB_NAME_TEMPLATE, mockDbName, DEV_ENV), result.DbName)

	result = loadDbConfig(STAGING_ENV)
	assert.Equal(fmt.Sprintf(DB_NAME_TEMPLATE, mockDbName, STAGING_ENV), result.DbName)

	result = loadDbConfig(TEST_ENV)
	assert.Equal(fmt.Sprintf(DB_NAME_TEMPLATE, mockDbName, TEST_ENV), result.DbName)
}
