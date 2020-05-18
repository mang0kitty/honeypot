package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mang0kitty/honeypot/honeypot"
	"github.com/mang0kitty/honeypot/state"
)

type API struct {
	Database *state.Database
	Honeypot *honey
}

func Handle(db *state.Database, h *honeypot.Honeypot) {
	r := mux.NewRouter()
	api := &API{
		Database: db,
		Honeypot: h,
	}
	r.HandleFunc("/stats", api.StatsHandler)

	log.Fatal(http.ListenAndServe(":3000", r))
}

func (api *API) StatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api)
}
