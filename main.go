package main

import (
	"github.com/RappyTV/PlayerRecordAPI/cron"
	"github.com/RappyTV/PlayerRecordAPI/db"
	"github.com/RappyTV/PlayerRecordAPI/env"
	"github.com/RappyTV/PlayerRecordAPI/router"
)

func main() {
	env.LoadEnv()
	db.Connect()
	defer db.Disconnect()
	cron.SetupScheduler()
	defer cron.ShutdownScheduler()
	router.SetupRouter()
}
