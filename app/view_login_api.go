package app

import (
	"strconv"

	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewLoginApi(sess session.Store, ctx *macaron.Context) {
	lar := &LoginApiResponse{}
	lar.SessionId = sess.ID()

	tgids := ctx.Params("telegramid")
	tgid, err := strconv.Atoi(tgids)
	if err != nil {
		loge(err)
	}

	err = sess.Set("tgid", int64(tgid))
	if err != nil {
		loge(err)
	}

	lar.TelegramId = int64(tgid)

	ctx.JSON(200, lar)
}

type LoginApiResponse struct {
	TelegramId int64  `json:"telegram_id"`
	SessionId  string `json:"session_id"`
}
