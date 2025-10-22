package router

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"player-record-api/db"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const Version = "1.0.0"

type VersionResponse struct {
	Version string `json:"version"`
}

type ServerResponse struct {
	Name         string `json:"server"`
	PlayerRecord int32  `json:"player_record"`
	Timestamp    int64  `json:"timestamp"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func SetupRouter() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(Recoverer)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Not found!"})
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Method not allowed!"})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(VersionResponse{Version: Version})
	})

	r.Get("/{serverName}", func(w http.ResponseWriter, r *http.Request) {
		serverName := chi.URLParam(r, "serverName")
		encoder := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")

		server, err := db.GetServer(serverName)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				w.WriteHeader(http.StatusNotFound)
				encoder.Encode(ErrorResponse{Error: "Server not found!"})
				return
			}

			panic(err)
		}

		encoder.Encode(&ServerResponse{
			Name:         server.Name,
			PlayerRecord: server.PlayerRecord,
			Timestamp:    server.Timestamp.Unix(),
		})
	})

	bind := os.Getenv("API_BIND")
	log.Println("Server listening on " + bind)
	log.Fatalln(http.ListenAndServe(bind, r))
}
