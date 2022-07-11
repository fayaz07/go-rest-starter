package repository

import (
	mappers "go-rest-starter/src/api/mappers"
	reqModels "go-rest-starter/src/api/models/req/auth"
)

type authRepo BaseRepository

var _authRepoInstance authRepo

func GetAuthRepo() authRepo {
	return _authRepoInstance
}

// --- implementation

func (a authRepo) Init(req reqModels.AuthInitReq) {
	_clientRepoInstance.Save(mappers.AuthInitReqToClientModel(req))
}
