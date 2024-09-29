package app

import (
	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewLoginApi(sess session.Store, ctx *macaron.Context) {
	lar := &LoginApiResponse{}
	lar.SessionId = sess.ID()

	tgids := ctx.Params("telegramid")
	err := sess.Set("tgid", tgids)
	if err != nil {
		loge(err)
	}

	lar.TelegramId = tgids

	ctx.JSON(200, lar)
}

type LoginApiResponse struct {
	TelegramId string `json:"telegram_id"`
	SessionId  string `json:"session_id"`
}
