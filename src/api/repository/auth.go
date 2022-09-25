package repository

import (
	reqModels "go-rest-starter/src/api/models/req"
)

type authRepo BaseRepository

var _authRepoInstance authRepo

func GetAuthRepo() authRepo {
	return _authRepoInstance
}

// --- implementation

func (a authRepo) InitSession(req reqModels.AuthInitReq) {
	deviceData, err := req.AuthInitReqToDeviceModel()
	if err != nil {
		_sessionRepoInstance.GetSession(req.IP, deviceData)
	}
}
