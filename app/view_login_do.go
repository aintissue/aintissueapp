package app

import (
	"log"

	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewloginDo(sess session.Store, ctx *macaron.Context) {
	// logs(prettyPrint(ctx.Data["user"]))
	// return "hello world"
	// tgid := ctx.Params("telegramid")
	sid := ctx.Params("sessionid")
	log.Println(sid)
	ctx.SetCookie("MacaronSession", sid)
	ctx.Redirect("/")
}
