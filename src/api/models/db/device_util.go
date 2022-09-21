package db

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
	return m.Type == MOBILE_CLIENT
}

func (m DeviceModel) IsDesktop() bool {
	return m.Type == DESKTOP_CLIENT
}

func (m DeviceModel) IsWeb() bool {
	return m.Type == WEB_CLIENT
}
