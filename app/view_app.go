package app

import (
	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewApp(sess session.Store, ctx *macaron.Context) {
	u := ctx.Data["User"].(*User)
	var chats []*Chat

	db.Find(&chats, &Chat{OwnerID: u.ID})
	ctx.Data["Chats"] = chats

	ctx.HTML(200, "home")
}
