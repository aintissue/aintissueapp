package app

import (
	"strconv"

	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewloginDo(sess session.Store, ctx *macaron.Context) {
	tgids := ctx.Params("telegramid")
	sid := ctx.Params("sessionid")
	ctx.SetCookie("MacaronSession", sid)

	tgid, err := strconv.Atoi(tgids)
	if err != nil {
		loge(err)
	}

	u := getUser(int64(tgid))
	if u.ID == 0 {
		ctx.Redirect("/login")
	}

	ctx.Redirect("/")
}
