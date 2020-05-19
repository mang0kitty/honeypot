package tests

import (
	"testing"

	"github.com/mang0kitty/honeypot/state"
)

func TestDatabase(t *testing.T) {
	database := state.NewDatabase()

	database.Add(&state.Record{
		RemoteAddr:  "127.0.0.1",
		Credentials: "admin:test",
		Protocol:    "ssh",
	})
	if len(database.Protocols) != 1 {
		t.Errorf("Record for SSH request wasn't added to the database.")
	}
}
