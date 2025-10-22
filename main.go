package main

import (
	"player-record-api/cron"
	"player-record-api/db"
	"player-record-api/env"
	"player-record-api/router"
)

func main() {
	env.LoadEnv()
	db.Connect()
	defer db.Disconnect()
	cron.SetupScheduler()
	defer cron.ShutdownScheduler()
	router.SetupRouter()
}
