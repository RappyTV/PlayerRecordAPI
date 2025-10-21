package main

import (
	"player-record-api/db"
	"player-record-api/env"
	"player-record-api/router"
)

// TODO: return schema kind of like this { "server": string, "player_record": int32, "timestamp": int64 }
// TODO: use https://laby.net/api/server/playercount/{server} to get graph data

func main() {
	env.LoadEnv()
	db.Connect()
	defer db.Disconnect()
	router.SetupRouter()
}
