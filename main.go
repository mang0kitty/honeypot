package main

import (
	"github.com/mang0kitty/honeypot/handlers"
	"github.com/mang0kitty/honeypot/honeypot"
	"github.com/mang0kitty/honeypot/protocols"
	"github.com/mang0kitty/honeypot/state"
)

func main() {
	database := state.NewDatabase()
	honeypot := &honeypot.Honeypot{Database: database, TotalVisits: 0}

	go func() {
		handlers.Handle(database, honeypot)

	}()

	protocols.Ssh(honeypot)

}
