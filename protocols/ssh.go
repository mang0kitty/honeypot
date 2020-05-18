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
			RemoteAddr:  ctx.RemoteAddr().String(),
			Credentials: fmt.Sprintf("%s:%s %x",ctx.User(),key.Type(), md5.Sum(key.Marshal()[len(key.Type()):])),
			Protocol: "ssh",
		}
		h.Database.Add(&record)
		log.Println(record)

		return false
	})

	passwordAuthOption := ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
		record := state.Record{
			RemoteAddr:  ctx.RemoteAddr().String(),
			Credentials: fmt.Sprintf("%s:%s",ctx.User(),pass),
			Protocol: "ssh",
		}

		h.Database.Add(&record)
		log.Println(record)
		log.Println(h.Database)
		return false
	})

	log.Println("**** Starting the SSH server on port 2222 ****")
	log.Fatal(ssh.ListenAndServe(":2222", nil, publicKeyOption, passwordAuthOption))
}