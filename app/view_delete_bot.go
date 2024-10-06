package app

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewDeleteBot(f *session.Flash, ctx *macaron.Context) {
	u := ctx.Data["User"].(*User)
	ids := ctx.Params("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		loge(err)
	}

	bot := getBot(uint(id))

	if bot.OwnerID != u.ID {
		f.Error("You are not the owner of this bot. This incident has been reported.")
		ctx.Redirect("/")
		return
	}

	cmd := exec.Command("docker", "stop", bot.Namespace)
	out, err := cmd.Output()
	if err != nil {
		loge(err)
		f.Error(err.Error())
		ctx.Redirect("/")
	}
	log.Println("Output: ", string(out))

	cmd = exec.Command("docker", "rm", bot.Namespace)
	out, err = cmd.Output()
	if err != nil {
		loge(err)
		f.Error(err.Error())
		ctx.Redirect("/")
	}
	log.Println("Output: ", string(out))

	cmd = exec.Command("rm", "-rf", fmt.Sprintf("data/bots/%s", bot.Dir))
	out, err = cmd.Output()
	if err != nil {
		loge(err)
		f.Error(err.Error())
		ctx.Redirect("/")
	}
	log.Println("Output: ", string(out))

	err = db.Delete(bot).Error
	if err != nil {
		loge(err)
		f.Error(err.Error())
		ctx.Redirect("/")
	}

	f.Success("Bot successfully deleted.")
	ctx.Redirect("/")
}
