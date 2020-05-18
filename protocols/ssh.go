package protocols

import(
	"crypto/md5"
	"fmt"
	"log"

	"github.com/gliderlabs/ssh"
	"github.com/mang0kitty/honeypot/state"
	"github.com/mang0kitty/honeypot/honeypot"
)

func Ssh(h *honeypot.Honeypot){

	ssh.Handle(func(s ssh.Session) {

	})
 
	publicKeyOption := ssh.PublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
		record := state.Record{
			User:        ctx.User(),
			RemoteAddr:  ctx.RemoteAddr().String(),
			Credentials: fmt.Sprintf("%s %x", key.Type(), md5.Sum(key.Marshal()[len(key.Type()):])),
		}
		h.TotalVisits = h.TotalVisits + 1
		h.Database.Add(&record)
		log.Println(record)

		return false
	})

	passwordAuthOption := ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
		record := state.Record{
			User:        ctx.User(),
			RemoteAddr:  ctx.RemoteAddr().String(),
			Credentials: pass,
		}

		h.Database.Add(&record)
		log.Println(record)
		log.Println(h.Database)
		return false
	})

	log.Println("**** Starting the SSH server on port 2222 ****")
	log.Fatal(ssh.ListenAndServe(":2222", nil, publicKeyOption, passwordAuthOption))
}