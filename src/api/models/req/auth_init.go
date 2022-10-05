package req

import (
	"go-rest-starter/src/api/models/consts"

	validation "github.com/go-ozzo/ozzo-validation"
)

type AuthInitReq struct {
	IP            string `json:"ip"`
	Device        string `json:"device" binding:"required"`
	OS            string `json:"os" binding:"required"`
	OSVersion     string `json:"osVersion" binding:"required"`
	ClientVersion string `json:"clientVersion" binding:"required"`

	// Fields for mobile
	Model      string `json:"model" binding:"required"`
	Maker      string `json:"maker" binding:"required"`
	Connection string `json:"connection" binding:"required"`

	// Fields for web
	UserAgent        string `json:"userAgent"`
	UserAgentVersion string `json:"userAgentVersion"`
	Host             string `json:"host"`
	Referer          string `json:"referer"`
}

func (a AuthInitReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.IP),
		validation.Field(&a.Device, validation.Required, validation.In(consts.DESKTOP_CLIENT, consts.MOBILE_CLIENT, consts.WEB_CLIENT)),
		validation.Field(&a.OS, validation.Required, validation.Length(3, 20)),
		validation.Field(&a.OSVersion, validation.Required, validation.Length(1, 12)),
		validation.Field(&a.ClientVersion, validation.Required, validation.Length(1, 12)),
		validation.Field(&a.Model, validation.Required, validation.Length(2, 24)),
		validation.Field(&a.Maker, validation.Required, validation.Length(2, 24)),
		validation.Field(&a.Connection, validation.Required, validation.Length(2, 24)),
		validation.Field(&a.UserAgent, validation.Length(2, 128)),
		validation.Field(&a.UserAgentVersion, validation.Length(1, 12)),
		validation.Field(&a.Host, validation.Length(2, 64)),
		validation.Field(&a.Referer, validation.Length(2, 64)),
	)
}
