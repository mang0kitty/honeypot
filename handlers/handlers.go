package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mang0kitty/honeypot/state"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type API struct {
	Database *state.Database
}

func Handle(db *state.Database) {
	r := mux.NewRouter()
	api := &API{
		Database: db,
	}
	r.HandleFunc("/stats", api.StatsHandler)
	r.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":3000", r))
}

func (api *API) StatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api) //immediately flushed to network, little stored in memory (buffer)
}
