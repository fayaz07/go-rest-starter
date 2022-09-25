package db

const (
	MOBILE_CLIENT  = "mobile"
	DESKTOP_CLIENT = "desktop"
	WEB_CLIENT     = "web"
)

func NewMobileDevice(
	os string,
	osVersion string,
	clientVersion string,
	maker string,
	model string,
) DeviceModel {
	return DeviceModel{
		Type:          MOBILE_CLIENT,
		OS:            os,
		OSVersion:     osVersion,
		ClientVersion: clientVersion,
		Maker:         maker,
		Model:         model,
	}
}

func NewWebDevice(
	os string,
	osVersion string,
	clientVersion string,
	maker string,
	model string,
	userAgent string,
	userAgentVersion string,
	referer string,
) DeviceModel {
	return DeviceModel{
		Type:             WEB_CLIENT,
		OS:               os,
		OSVersion:        osVersion,
		ClientVersion:    clientVersion,
		Maker:            maker,
		Model:            model,
		UserAgent:        userAgent,
		UserAgentVersion: userAgentVersion,
		Referer:          referer,
	}
}

func NewDesktopDevice(
	os string,
	osVersion string,
	clientVersion string,
	maker string,
	model string,
) DeviceModel {
	return DeviceModel{
		Type:          DESKTOP_CLIENT,
		OS:            os,
		OSVersion:     osVersion,
		ClientVersion: clientVersion,
		Maker:         maker,
		Model:         model,
	}
}

func (m DeviceModel) IsMobile() bool {
	return IsMobile(m.Type)
}

func IsMobile(s string) bool {
	return s == MOBILE_CLIENT
}

func (m DeviceModel) IsDesktop() bool {
	return IsDesktop(m.Type)
}

func IsDesktop(s string) bool {
	return s == DESKTOP_CLIENT
}

func (m DeviceModel) IsWeb() bool {
	return IsWeb(m.Type)
}

func IsWeb(s string) bool {
	return s == WEB_CLIENT
}

func (m DeviceModel) Print() {
	PrettyPrint(m)
}
