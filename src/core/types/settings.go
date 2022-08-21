package types

type AppSettings struct {
	AppName string

	PublicDir        string
	UploadsPublicDir string
	StaticPublicDir  string
	LogsDir          string
	DotEnvsDir       string
	DotEnvFile       string

	KeysDir            string
	AccessTokenPvtKey  string
	AccessTokenPubKey  string
	RefreshTokenPvtKey string
	RefreshTokenPubKey string
}
