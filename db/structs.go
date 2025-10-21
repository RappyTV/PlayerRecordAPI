package db

import "time"

type Server struct {
	Name         string    `bson:"name" json:"name"`
	PlayerRecord int32     `bson:"player_record" json:"player_record"`
	Timestamp    time.Time `bson:"timestamp" json:"timestamp"`
}
