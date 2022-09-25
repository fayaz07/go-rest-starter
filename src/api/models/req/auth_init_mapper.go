package req

import (
	"errors"
	dbModels "go-rest-starter/src/api/models/db"
)

func (data AuthInitReq) AuthInitReqToDeviceModel() (dbModels.DeviceModel, error) {
	switch data.Device {
	case dbModels.MOBILE_CLIENT:
		return data.buildMobileDevice(), nil
	case dbModels.DESKTOP_CLIENT:
		return data.buildDesktopDevice(), nil
	case dbModels.WEB_CLIENT:
		return data.buildWebDevice(), nil
	default:
		return dbModels.DeviceModel{}, errors.New("unsupported device type")
	}
}

func (data AuthInitReq) buildMobileDevice() dbModels.DeviceModel {
	return dbModels.NewMobileDevice(data.OS, data.OSVersion,
		data.ClientVersion, data.Manufacturer, data.Model)
}

func (data AuthInitReq) buildDesktopDevice() dbModels.DeviceModel {
	return dbModels.NewDesktopDevice(data.OS, data.OSVersion,
		data.ClientVersion, data.Manufacturer, data.Model)
}

func (data AuthInitReq) buildWebDevice() dbModels.DeviceModel {
	return dbModels.NewWebDevice(data.OS, data.OSVersion,
		data.ClientVersion, data.Manufacturer, data.Model,
		data.UserAgent, data.UserAgentVersion, data.Referer)
}
