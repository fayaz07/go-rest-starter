package repository

import (
	modelRef "go-rest-starter/src/api/models/db"
	log "go-rest-starter/src/core/logger"

	"go.mongodb.org/mongo-driver/bson"
)

type clientRepo BaseRepository

var _clientRepoInstance clientRepo

func GetClientRepo() clientRepo {
	return _clientRepoInstance
}

func initClientColln() {
	log.I("Initializing client collection in repository")
	_clientRepoInstance = clientRepo{
		Name:  "ClientRepository",
		model: modelRef.GetClientCollection(db),
	}
}

// --- implementation

func (c clientRepo) Save(data modelRef.ClientModel) (modelRef.ClientModel, error) {
	c.GetByIP(data.IP)
	ins, err := c.model.InsertOne(GetDbCtx(), data)
	if err != nil {
		return modelRef.ClientModel{}, err
	}
	log.If("%v", ins)
	return modelRef.ClientModel{}, nil
}

func (c clientRepo) GetByIP(ip string) {
	var result modelRef.ClientModel
	err := c.model.FindOne(GetDbCtx(), bson.M{"ip": ip}).Decode(&result)
	if err == nil {
		result.Print()
	} else {
		log.E(err.Error())
	}
}
