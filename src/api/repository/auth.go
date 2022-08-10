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

func (a authRepo) Init(req reqModels.AuthInitReq) {
	_clientRepoInstance.Save(req.AuthInitReqToClientModel())
}
