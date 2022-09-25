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

func (a authRepo) InitSession(req reqModels.AuthInitReq) (string, error) {

	session, err := _sessionRepoInstance.GetSession(req.IP)
	if err != nil {
		return "", err
	}

	deviceData, err := req.AuthInitReqToDeviceModel(session.ID)
	if err != nil {
		return "", err
	}
	_deviceRepoInstance.GetDevice(deviceData)
	return "done", nil
}
