package mappers

import (
	model "go-rest-starter/src/api/models/db"
	reqModels "go-rest-starter/src/api/models/req/auth"
)

func AuthInitReqToClientModel(data reqModels.AuthInitReq) model.ClientModel {
	return model.ClientModel{
		IP:      data.IP,
		Device:  data.Device,
		OS:      data.OS,
		Version: data.Version,

		Model:        data.Model,
		Manufacturer: data.Manufacturer,
		Connection:   data.Connection,

		UserAgent: data.UserAgent,
		Host:      data.Host,
		Browser:   data.Browser,

		Referer: data.Referer,
	}
}
