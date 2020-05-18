package state

import (
	"encoding/json"
	"fmt"
)

type Database struct {
	TotalVisits int                  `json:"totalVisits"`
	Protocols   map[string]*Protocol `json:"protocols"`
}

func NewDatabase() *Database {
	return &Database{
		TotalVisits: 0,
		Protocols:   map[string]*Protocol{},
	}
}

func (db *Database) Add(record *Record) {
	db.TotalVisits++

	if p, ok := db.Protocols[record.Protocol]; ok {
		p.Visits++

		p.RemoteAddr[record.RemoteAddr]++
		p.Credentials[record.Credentials]++
	} else {
		protocol := Protocol{
			Name:        record.Protocol,
			Visits:      1,
			RemoteAddr:  map[string]int{record.RemoteAddr: 1},
			Credentials: map[string]int{record.Credentials: 1},
		}
		db.Protocols[record.Protocol] = &protocol
	}

	fmt.Println(db.String())
}

func (db *Database) String() string {
	b, _ := json.Marshal(db)
	return string(b)
}
