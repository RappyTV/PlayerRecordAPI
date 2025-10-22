package labynet

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"player-record-api/router"
)

const endpoint string = "https://laby.net/api/server/playercount/%v"

type PlayerCountGraphPoint struct {
	Timestamp   string `bson:"x" json:"x"`
	PlayerCount int32  `bson:"y" json:"y"`
}

func GetServerPlayerRecord(server string) (*PlayerCountGraphPoint, error) {
	request, _ := http.NewRequest("GET", fmt.Sprintf(endpoint, server), nil)
	request.Header.Add("User-Agent", "PlayerRecordAPI v"+router.Version+" (github.com/RappyTV/PlayerRecordAPI)")
	res, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		resBody, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("failed to fetch player record for server \"%v\" with status code %v: %v", server, res.StatusCode, resBody)
	}

	var graphData []PlayerCountGraphPoint

	err = json.NewDecoder(res.Body).Decode(&graphData)

	if err != nil {
		return nil, err
	}

	if len(graphData) == 0 {
		return nil, fmt.Errorf("no graph data returned for server %v", server)
	}

	var playerRecord *PlayerCountGraphPoint = &graphData[0]
	for _, point := range graphData[1:] {
		if point.PlayerCount > playerRecord.PlayerCount {
			playerRecord = &point
		}
	}

	return playerRecord, nil
}
