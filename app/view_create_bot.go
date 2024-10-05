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
	u := ctx.Data["User"].(*User)
	if len(u.getBots()) > 0 {
		f.Error("You can only have one bot in trial, please upgrade to create unlimited number of bots.")
		ctx.Redirect("/create")
		return
	}

	if len(botf.TelegramKey) == 0 || len(botf.Name) == 0 {
		f.Error("Both fields are required is required.")
		ctx.Redirect("/create")
		return
	}

	cmd := exec.Command("cp", "-R", "data/aintissuebot", fmt.Sprintf("data/bots/bot%d", getBotCount()))
	out, err := cmd.Output()
	if err != nil {
		loge(err)
		f.Error(err.Error())
		ctx.Redirect("/create")
		return
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
		return
	}

	log.Println("Output: ", string(out))

	botName := fmt.Sprintf("bot%d-bot-1", getBotCount())
	bot := &Bot{
		Name:      botf.Name,
		Namespace: botName,
		OwnerID:   ctx.Data["User"].(*User).ID,
	}

	err = db.Save(bot).Error
	if err != nil {
		loge(err)
		f.Error(err.Error())
		ctx.Redirect("/create")
		return
	}

	increaseBotCount()

	f.Success("Bot successfully created.")
	ctx.Redirect("/")
}

type BotForm struct {
	TelegramKey string `form:"telegram_key"`
	Name        string `form:"name"`
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
