package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewCreateBot(sess session.Store, ctx *macaron.Context) {
	ctx.HTML(200, "bot")
}

func viewDoCreateBot(botf BotForm, f *session.Flash, ctx *macaron.Context) {
	cmd := exec.Command("cp", "-R", "data/aintissuebot", fmt.Sprintf("data/bots/bot%d", getBotCount()))
	out, err := cmd.Output()
	if err != nil {
		loge(err)
		f.Error(err.Error())
		ctx.Redirect("/create")
	}

	log.Println("Output: ", string(out))

	replaceInFile(fmt.Sprintf("data/bots/bot%d/data/config.yaml", getBotCount()), "TELEGRAM_BOT_API_KEY", botf.TelegramKey)

	cmd = exec.Command("docker", "compose", "up", "-d")
	cmd.Dir = fmt.Sprintf("data/bots/bot%d", getBotCount())

	out, err = cmd.Output()
	if err != nil {
		loge(err)
		f.Error(err.Error())
		ctx.Redirect("/create")
	}

	log.Println("Output: ", string(out))

	botName := fmt.Sprintf("bot%d-bot-1", getBotCount())
	bot := &Bot{
		Name:      botName,
		Namespace: botName,
		OwnerID:   ctx.Data["User"].(*User).ID,
	}

	err = db.Save(bot).Error
	if err != nil {
		loge(err)
		f.Error(err.Error())
		ctx.Redirect("/create")
	}

	increaseBotCount()

	f.Success("Bot successfully created.")
	ctx.Redirect("/")
}

type BotForm struct {
	TelegramKey string `form:"telegram_key"`
}

func replaceInFile(path string, old string, new string) {
	f, err := os.ReadFile(path)
	if err != nil {
		loge(err)
	}

	fs := string(f)
	fs = strings.ReplaceAll(fs, old, new)
	f = []byte(fs)
	err = os.WriteFile(path, f, 0o664)
	if err != nil {
		loge(err)
	}
}
