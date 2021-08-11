package main

import (
	"fmt"
	core "go-rest-starter/src/core"
)

func main() {

	appConfig := core.GetAppConfig()
	fmt.Println("#v", appConfig)

	core.GetDbConnection(appConfig.DB)
}
