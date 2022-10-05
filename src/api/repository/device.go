package repository

import (
	modelRef "go-rest-starter/src/api/models/db"
	log "go-rest-starter/src/core/logger"
	"go-rest-starter/src/utils/helpers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (c deviceRepo) GetDeviceId(data modelRef.DeviceModel) (primitive.ObjectID, error) {
	exists, err := c.SearchDevice(data)
	if err == nil {
		return exists.ID, nil
	}
	ctx := helpers.GetDbRWContext()
	ins, err := c.model.InsertOne(ctx.Ctx, data)
	ctx.Cancel()
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return id(ins)
}

func (c deviceRepo) SearchDevice(data modelRef.DeviceModel) (*modelRef.DeviceModel, error) {
	ctx := helpers.GetDbRWContext()
	var result *modelRef.DeviceModel
	err := c.model.FindOne(ctx.Ctx, searchBson(data)).Decode(&result)
	ctx.Cancel()
	if err == nil {
		return result, nil
	}
	return nil, err
}

func searchBson(data modelRef.DeviceModel) bson.M {
	return bson.M{
		"cV":    data.ClientVersion,
		"_sId":  data.SessionID,
		"_uId":  data.UserID,
		"type":  data.Type,
		"maker": data.Maker,
		"model": data.Model,
		"os":    data.OS,
	}
}
