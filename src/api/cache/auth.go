package cache

import (
	"encoding/json"
	"fmt"
	"go-rest-starter/src/core/logger"
	utils "go-rest-starter/src/utils/helpers"
	"time"

	uuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const FIVE_MINUTES = time.Minute * 5

type AuthExchangeCache struct {
	Token string             `json:"t" binding:"required"`
	SId   primitive.ObjectID `json:"s" binding:"required"`
	DId   primitive.ObjectID `json:"d" binding:"required"`
}

func (i AuthExchangeCache) Print() {
	logger.I(utils.Pretty(i))
}

func (i AuthExchangeCache) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func UnMarshalToAuthExchangeCache(data string) (AuthExchangeCache, error) {
	var ae AuthExchangeCache
	err := json.Unmarshal([]byte(data), &ae)
	if err != nil {
		return AuthExchangeCache{}, err
	}
	return ae, nil
}

func getAuthExchangeObject(token string, sId primitive.ObjectID, dId primitive.ObjectID) ([]byte, error) {
	return AuthExchangeCache{SId: sId, DId: dId, Token: token}.MarshalBinary()
}

func getIPExchKey(ip string) string {
	return fmt.Sprintf("exch_%s", ip)
}

func getNewUUiD() string {
	return uuid.New().String()
}

func ValidateAuthExchangeTokenByIP() bool {
	return false
}

func CreateAuthExchangeToken(sId primitive.ObjectID, dId primitive.ObjectID, ip string) (string, error) {
	ctx := utils.GetCacheRWContext()
	token := getNewUUiD()
	data, err := getAuthExchangeObject(token, sId, dId)
	if err != nil {
		return "", err
	}
	err = client.Set(ctx.Ctx, getIPExchKey(ip), data, FIVE_MINUTES).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetAuthExchToken(ip string) (AuthExchangeCache, bool) {
	ctx := utils.GetCacheRWContext()
	res := client.Get(ctx.Ctx, getIPExchKey(ip))
	ctx.Cancel()
	if res.Err() != nil {
		return AuthExchangeCache{}, false
	}
	obj, err := UnMarshalToAuthExchangeCache(res.Val())
	if err != nil {
		return AuthExchangeCache{}, false
	}
	return obj, true
}
