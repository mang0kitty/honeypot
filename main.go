package main

import (
	"github.com/mang0kitty/honeypot/honeypot"
	"github.com/mang0kitty/honeypot/state"
)

func main() {
	database := state.NewDatabase()
	honeypot := honeypot.Honeypot{Database: database, TotalVisits: 0}
	honeypot.Run()
}
