package main

import (
	"money-api-transfer/api/config"
	"money-api-transfer/cmd/server"
)

func main() {
	appConf := config.GetConfig()

	server.StartAPI(appConf)

}
