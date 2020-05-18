package state

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/mang0kitty/honeypot/helpers"
	"github.com/mang0kitty/honeypot/profile"
)

type Database struct {
	Users map[string]*profile.User `json:"users"`
}

func NewDatabase() *Database {
	return &Database{
		Users: map[string]*profile.User{},
	}
}

func (db *Database) Add(record *Record) {

	if usr, ok := db.Users[record.RemoteAddr]; ok {
		if !helpers.Contains(usr.Usernames, record.User) {
			usr.Usernames = append(usr.Usernames, record.User)
		}

		if !helpers.Contains(usr.Credentials, record.Credentials) {
			usr.Credentials = append(usr.Credentials, record.Credentials)
		}

		usr.Visits++
	} else {
		user := profile.User{
			Usernames:   []string{record.User},
			Credentials: []string{record.Credentials},
			RemoteAddr:  record.RemoteAddr,
			Visits:      1,
		}

		db.Users[record.RemoteAddr] = &user
	}

	fmt.Println(db.String())
}

func (db *Database) String() string {
	b := bytes.NewBuffer([]byte{})
	json.NewEncoder(b).Encode(db)
	return b.String()
}
