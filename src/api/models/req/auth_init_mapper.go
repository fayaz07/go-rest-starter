package req

import (
	"errors"
	dbModels "go-rest-starter/src/api/models/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (data AuthInitReq) AuthInitReqToDeviceModel(sId primitive.ObjectID) (dbModels.DeviceModel, error) {

	if !dbModels.IsValidDevice(data.Device) {
		return dbModels.DeviceModel{}, errors.New("unsupported device type")
	}

	return dbModels.DeviceModel{
		SessionID:        sId,
		Type:             data.Device,
		Connection:       data.Connection,
		OS:               data.OS,
		OSVersion:        data.OSVersion,
		ClientVersion:    data.ClientVersion,
		Maker:            data.Maker,
		Model:            data.Model,
		UserAgent:        data.UserAgent,
		UserAgentVersion: data.UserAgentVersion,
		Referer:          data.Referer,
	}, nil
}
