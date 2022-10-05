package db

type AuthRole string

const (
	ADMIN AuthRole = "admin"
)

type AuthProvider string

const (
	EMAIL AuthProvider = "email"
)

func NewAdmin(email string, password string) AuthModel {
	return AuthModel{
		Email:    email,
		Password: password,
		Role:     ADMIN,
		Provider: EMAIL,
	}
}
