package db

import (
	"encoding/json"
	log "go-rest-starter/src/core/logger"
)

func PrettyPrint(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Ef("error: %v", err)
	}
	log.I(string(b))
}
