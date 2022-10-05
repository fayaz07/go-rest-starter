package wrappers

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var r = GetResponseWrapper()

func TestResponseWrapper_Success(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(gin.H{"status": "success", "message": "User login success"}, r.Success("User login success", nil))
	assert.NotEqual(gin.H{"status": "failed", "message": "User login success"}, r.Success("User login success", nil))
}

func TestResponseWrapper_Failed(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(gin.H{"status": "failed", "message": "User login failed"}, r.Error("User login failed", nil))
	assert.NotEqual(gin.H{"status": "success", "message": "User login failed"}, r.Error("User login failed", nil))
}

func TestResponseWrapper_SuccessWithRawData(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(gin.H{"status": "success", "message": "User login success", "data": 10}, r.Success("User login success", 10))
	assert.Equal(gin.H{"status": "success", "message": "User login success", "data": "SUCCESS"}, r.Success("User login success", "SUCCESS"))
	assert.Equal(gin.H{"status": "success", "message": "User login success", "data": 10.3}, r.Success("User login success", 10.3))

	assert.NotEqual(gin.H{"status": "failed", "message": "User login success"}, r.Success("User login success", 10))
	assert.NotEqual(gin.H{"status": "failed", "message": "User login success", "data": nil}, r.Success("User login success", 10))
}

func TestResponseWrapper_SuccessWithTypedData(t *testing.T) {
	assert := assert.New(t)

	type user struct {
		name     string
		id       int
		salary   float32
		verified bool
	}

	assert.Equal(
		gin.H{
			"status":  "success",
			"message": "User login success",
			"data":    user{name: "john", id: 12, salary: 343.2, verified: true},
		},
		r.Success(
			"User login success",
			user{name: "john", id: 12, salary: 343.2, verified: true},
		),
	)
	assert.NotEqual(
		gin.H{
			"status":  "success",
			"message": "User login success",
			"data":    gin.H{"name": "john", "id": 12, "salary": 343.2, "verified": true},
		},
		r.Success(
			"User login success",
			user{name: "john", id: 12, salary: 343.2, verified: true},
		),
	)
}

func TestResponseWrapper_FailedWithRawError(t *testing.T) {
	assert := assert.New(t)

	type userError struct {
		email    string
		password string
	}

	assert.Equal(
		gin.H{"status": "failed", "message": "User login failed", "error": userError{email: "Invalid email", password: "Password is not strong"}},
		r.Error("User login failed", userError{email: "Invalid email", password: "Password is not strong"}),
	)
	assert.NotEqual(
		gin.H{"status": "success", "message": "User login failed", "data": userError{email: "Invalid email", password: "Password is not strong"}},
		r.Error("User login failed", userError{email: "Invalid email", password: "Password is not strong"}),
	)
}
