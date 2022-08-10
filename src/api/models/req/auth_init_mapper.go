package req

import (
	dbModels "go-rest-starter/src/api/models/db"
)

func (data AuthInitReq) AuthInitReqToClientModel() dbModels.ClientModel {
	return dbModels.ClientModel{
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
