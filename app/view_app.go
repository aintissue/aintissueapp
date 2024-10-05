package app

import (
	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewApp(sess session.Store, ctx *macaron.Context) {
	u := ctx.Data["User"].(*User)
	ctx.Data["Upgrade"] = 0

	ctx.Data["Chats"] = u.getChats()
	ctx.Data["Bots"] = u.getBots()

	if u.Plan == PlanFree && (len(u.getChats()) > 0 || len(u.getBots()) > 0) {
		ctx.Data["Upgrade"] = 1
	}

	ctx.HTML(200, "home")
}
