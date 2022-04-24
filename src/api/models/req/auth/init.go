package auth

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type AuthInitReq struct {
	Device  string `json:"device" binding:"required"`
	OS      string `json:"os" binding:"required"`
	Version string `json:"version"`

	// Fields for mobile
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	Connection   string `json:"connection"`

	// Fields for web
	UserAgent string `json:"user_agent"`
	Host      string `json:"host"`
	Browser   string `json:"browser"`
}

// {
//   "device": "mobile",
//   "os": "android",
//   "version": "7.0",
//   "model": "Samsung Galaxy S7",
//   "manufacturer": "Samsung",
//   "connection": "wifi"
// }

// {
//   "host": "localhost",
//   "device": "windows",
//   "os": "windows",
//   "browser": "chrome",
//   "browser_version": "80.0.3987.149"
// }

const Mobile_key = "mobile"
const Desktop_key = "desktop"

func (a AuthInitReq) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Device, validation.Required, validation.In(Mobile_key, Desktop_key)),
		validation.Field(&a.OS, validation.Required, validation.Length(3, 20)),
		validation.Field(&a.Version, validation.Required, validation.Length(2, 24)),
		validation.Field(&a.Model, validation.Length(2, 24)),
		validation.Field(&a.Manufacturer, validation.Length(2, 24)),
		validation.Field(&a.Connection, validation.Length(2, 24)),
		validation.Field(&a.UserAgent, validation.Length(2, 24)),
		validation.Field(&a.Host, validation.Length(2, 24)),
		validation.Field(&a.Browser, validation.Length(2, 24)),
	)
}
