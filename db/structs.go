package db

import "time"

type Server struct {
	Name         string    `bson:"name"`
	PlayerRecord int32     `bson:"player_record"`
	Timestamp    time.Time `bson:"timestamp"`
}
