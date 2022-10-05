package db

import (
	"go-rest-starter/src/api/models/consts"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsValidDevice(d string) bool {
	return d == consts.MOBILE_CLIENT || d == consts.DESKTOP_CLIENT || d == consts.WEB_CLIENT
}

func NewMobileDevice(
	os string,
	osVersion string,
	clientVersion string,
	maker string,
	model string,
	sId primitive.ObjectID,
) DeviceModel {
	return DeviceModel{
		SessionID:     sId,
		Type:          consts.MOBILE_CLIENT,
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
	sId primitive.ObjectID,
) DeviceModel {
	return DeviceModel{
		SessionID:        sId,
		Type:             consts.WEB_CLIENT,
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
	sId primitive.ObjectID,
) DeviceModel {
	return DeviceModel{
		SessionID:     sId,
		Type:          consts.DESKTOP_CLIENT,
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
	return s == consts.MOBILE_CLIENT
}

func (m DeviceModel) IsDesktop() bool {
	return IsDesktop(m.Type)
}

func IsDesktop(s string) bool {
	return s == consts.DESKTOP_CLIENT
}

func (m DeviceModel) IsWeb() bool {
	return IsWeb(m.Type)
}

func IsWeb(s string) bool {
	return s == consts.WEB_CLIENT
}

func (m DeviceModel) Print() {
	PrettyPrint(m)
}
