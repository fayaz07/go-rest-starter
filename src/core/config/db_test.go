package config

import (
	"fmt"
	"go-rest-starter/src/utils/constants"
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
	t.Setenv(constants.DB_NAME, mockDbName)
	t.Setenv(constants.DB_HOST, mockDbHost)
	t.Setenv(constants.DB_PORT, mockDbPort)
	t.Setenv(constants.DB_USERNAME, mockDbUsername)
	t.Setenv(constants.DB_PASSWORD, mockDbPassword)
}

func TestLoadDbConfig(t *testing.T) {
	assert := assert.New(t)

	_setDbConfigMocks(t)

	result := loadDbConfig(constants.PROD_ENV)

	assert.Equal(mockDbName, result.DbName)
	assert.Equal(mockDbHost, result.DbHost)
	assert.Equal(mockDbPort, result.DbPort)
	assert.Equal(mockDbUsername, result.DbUsername)
	assert.Equal(mockDbPassword, result.DbPassword)
}

func TestLoadDbConfigDbNames(t *testing.T) {
	assert := assert.New(t)

	_setDbConfigMocks(t)

	result := loadDbConfig(constants.PROD_ENV)
	assert.Equal(mockDbName, result.DbName)

	result = loadDbConfig(constants.DEV_ENV)
	assert.Equal(fmt.Sprintf(DB_NAME_TEMPLATE, mockDbName, constants.DEV_ENV), result.DbName)

	result = loadDbConfig(constants.STAGING_ENV)
	assert.Equal(fmt.Sprintf(DB_NAME_TEMPLATE, mockDbName, constants.STAGING_ENV), result.DbName)

	result = loadDbConfig(constants.TEST_ENV)
	assert.Equal(fmt.Sprintf(DB_NAME_TEMPLATE, mockDbName, constants.TEST_ENV), result.DbName)
}
