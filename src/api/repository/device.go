package repository

import (
	modelRef "go-rest-starter/src/api/models/db"
	log "go-rest-starter/src/core/logger"

	"go.mongodb.org/mongo-driver/bson"
)

type deviceRepo BaseRepository

var _deviceRepoInstance deviceRepo

func GetDeviceRepo() deviceRepo {
	return _deviceRepoInstance
}

func initDeviceColln() {
	log.I("Initializing device collection in repository")
	_deviceRepoInstance = deviceRepo{
		Name:  "deviceRepository",
		model: modelRef.GetDeviceCollection(db),
	}
}

// --- implementation

func (c deviceRepo) GetDeviceModel(data modelRef.DeviceModel) (modelRef.DeviceModel, error) {
	// doc, err := c.GetByIP(data.Maker)
	ins, err := c.model.InsertOne(GetDbCtx(), data)
	if err != nil {
		return modelRef.DeviceModel{}, err
	}
	log.If("%v", ins)
	return modelRef.DeviceModel{}, nil
}

func (c deviceRepo) GetByIP(ip string) (*modelRef.DeviceModel, error) {
	var result *modelRef.DeviceModel
	err := c.model.FindOne(GetDbCtx(), bson.M{"ip": ip}).Decode(&result)
	if err == nil {
		return result, nil
	}
	return nil, err
}
