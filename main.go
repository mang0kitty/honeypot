package main

import (
	"github.com/mang0kitty/honeypot/handlers"
	"github.com/mang0kitty/honeypot/honeypot"
	"github.com/mang0kitty/honeypot/protocols"
	"github.com/mang0kitty/honeypot/state"
)

func main() {
	database := state.NewDatabase()
	honeypot := &honeypot.Honeypot{Database: database}

	go func() {
		handlers.Handle(database)
	}()

	protocols.Ssh(honeypot)

}
