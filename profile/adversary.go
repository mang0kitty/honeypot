package profile

type User struct {
	Usernames   []string `json:"usernames"`
	Credentials []string `json:"credentials"`
	RemoteAddr  string   `json:"addr"`
	Visits      int      `json:"visits"`
}
