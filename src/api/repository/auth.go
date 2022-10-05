package repository

import (
	"go-rest-starter/src/api/cache"
	reqModels "go-rest-starter/src/api/models/req"
)

type authRepo BaseRepository

var _authRepoInstance authRepo

func GetAuthRepo() authRepo {
	return _authRepoInstance
}

// --- implementation
/*
	- client sends all the required details about them
	- server checks if there is a session instance available with that IP
		-	if found a session, then server checks for deviceId by matching sessionId,
			device details and returns deviceId (either fetching the existing or creating a new one)
		- else server creates a new session, creates a new device id
	- once server is ready with sessionId and deviceId, it creates an exchange auth token.
		-	client makes use of that auth exchange token to either login/register.
			This token is valid only for 5 minutes
*/
func (a authRepo) InitSession(req reqModels.AuthInitReq) (string, error) {
	session, err := _sessionRepoInstance.GetSession(req.IP)
	if err != nil {
		return "", err
	}

	deviceData, err := req.AuthInitReqToDeviceModel(session.ID)
	if err != nil {
		return "", err
	}
	deviceId, err := _deviceRepoInstance.GetDeviceId(deviceData)

	if err != nil {
		return "", err
	}

	token, err := cache.CreateAuthExchangeToken(session.ID, deviceId, session.IP)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (a authRepo) AuthExchangeTokenExists(ip string) bool {
	_, exists := cache.GetAuthExchToken(ip)
	return exists
}
