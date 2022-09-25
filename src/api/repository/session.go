package repository

import (
	models "go-rest-starter/src/api/models/db"
	log "go-rest-starter/src/core/logger"

	"go.mongodb.org/mongo-driver/bson"
)

type sessionRepo BaseRepository

var _sessionRepoInstance sessionRepo

func GetSessionRepo() sessionRepo {
	return _sessionRepoInstance
}

func initSessionCollection() {
	log.I("Initializing device collection in repository")
	_deviceRepoInstance = deviceRepo{
		Name:  "sessionRepo",
		model: models.GetSessionCollection(db),
	}
}

// --- implementation

func (r sessionRepo) GetSession(ip string, device models.DeviceModel) (*models.SessionModel, error) {
	doc, err := r.GetByIP(ip)
	if err == nil {
		return doc, nil
	}

	return &models.SessionModel{}, nil
}

func (r sessionRepo) createSession(ip string) (*models.SessionModel, error) {
	ctx := GetRWCtx()
	data := models.SessionModel{IP: ip}
	r.model.InsertOne(ctx.Ctx, data)
}

func (r sessionRepo) GetByIP(ip string) (*models.SessionModel, error) {
	var result *models.SessionModel
	err := r.model.FindOne(GetDbCtx(), bson.M{"ip": ip}).Decode(&result)
	if err == nil {
		return result, nil
	}
	return nil, err
}
