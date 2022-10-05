package types

// JwtConfig - configuration of JWT keys
type JwtConfig struct {
	AccessTokenExpiry  int
	RefreshTokenExpiry int
}
