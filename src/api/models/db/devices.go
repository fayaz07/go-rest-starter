package db

const (
	MOBILE_CLIENT  = "mobile"
	DESKTOP_CLIENT = "desktop"
	BROWSER_CLIENT = "browser"
)

type DeviceModel struct {
	Type          string `json:"type" binding:"required" bson:"type"`
	OS            string `json:"os" binding:"required" bson:"os"`
	OSVersion     string `json:"osV" binding:"required" bson:"osV"`
	ClientVersion string `json:"cV" binding:"required" bson:"cV"`

	Maker string `json:"maker" binding:"required" bson:"maker"`
	Model string `json:"model" binding:"required" bson:"model"`

	UserAgent        string `json:"ua" bson:"ua"`
	UserAgentVersion string `json:"uaV" bson:"uaV"`
}
