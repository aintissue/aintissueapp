package app

import (
	"log"
	"os/exec"

	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewStartBot(sess session.Store, ctx *macaron.Context) string {
	cmd := exec.Command("docker", "compose", "up", "-d")
	cmd.Dir = "./bot1/"

	out, err := cmd.Output()
	if err != nil {
		loge(err)
	}

	log.Println("Output: ", string(out))

	return "Hello world"
}
