package app

import (
	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewLogout(sess session.Store, ctx *macaron.Context) {
	err := sess.Delete("tgid")
	if err != nil {
		loge(err)
	}
	ctx.Redirect("/")
}
