package state

/*
The Protocol represents a protocol used to request access to the honeypot
Name is the name of the protocol
Visits is the total number of requests made with the protocol
RemoteAddr is the IP address of the adversary requesting access
Credentials map the public key/passwords used to the number of times they are used

*/

type Protocol struct {
	Name        string         `json:"name"`
	Visits      int            `json:"visits`
	RemoteAddr  map[string]int `json:"addresses`
	Credentials map[string]int `json:"credentials"`
}
