package app

import (
	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewApp(sess session.Store, ctx *macaron.Context) {
	// logs(prettyPrint(ctx.Data["user"]))
	// return "hello world"
	ctx.HTML(200, "home")
}
