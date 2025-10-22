package cron

import (
	"log"
	"os"
	"player-record-api/db"
	"player-record-api/labynet"
	"strings"
	"time"

	"github.com/go-co-op/gocron/v2"
)

var s gocron.Scheduler

func SetupScheduler() {
	_s, err := gocron.NewScheduler(gocron.WithLocation(time.FixedZone(os.Getenv("TZ"), 0)))

	if err != nil {
		log.Fatalln("Failed to create cron scheduler", err)
	}

	_, err = _s.NewJob(gocron.CronJob(os.Getenv("CRON_SCHEDULE"), false), gocron.NewTask(task))

	if err != nil {
		log.Fatalln("Failed to create cron job", err)
	}

	_s.Start()
	s = _s
}

func ShutdownScheduler() error {
	return s.Shutdown()
}

func task() {
	var servers []string = strings.Split(os.Getenv("FETCHED_SERVERS"), ",")

	for _, server := range servers {
		playerRecord, err := labynet.GetServerPlayerRecord(server)

		if err != nil {
			log.Printf("Failed to fetch player record for server %v: %v\n", server, err)
			return
		}

		currentRecord, err := db.GetServer(server)

		if err != nil {
			log.Println("Failed to get current server data from database:", err)
			return
		}

		if playerRecord.PlayerCount > currentRecord.PlayerRecord {
			currentRecord.PlayerRecord = playerRecord.PlayerCount
			parsedTime, err := time.Parse("2006-01-02 15:04:05", playerRecord.Timestamp)

			if err != nil {
				log.Println("Failed to parse timestamp:", err)
				return
			}

			currentRecord.Timestamp = parsedTime

			err = db.UpdateServer(currentRecord)

			if err != nil {
				log.Println("Failed to update server data in database:", err)
				return
			}

			log.Printf("Updated player record for server %v to %v players at timestamp %v\n", server, playerRecord.PlayerCount, playerRecord.Timestamp)
		}
	}
}
