package app

import (
	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewCreateBot(sess session.Store, ctx *macaron.Context) {
	ctx.HTML(200, "bot")
}

func viewDoCreateBot(sess session.Store, ctx *macaron.Context) string {
	// cmd := exec.Command("docker", "compose", "up", "-d")
	// cmd.Dir = "./bot1/"

	// out, err := cmd.Output()
	// if err != nil {
	// 	loge(err)
	// }

	// log.Println("Output: ", string(out))

	return "Hello world"
}

type BotForm struct {
	TelegramKey string `form:"telegram_key"`
}
