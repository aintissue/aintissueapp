package app

import (
	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewApp(sess session.Store, ctx *macaron.Context) string {
	// logs(sess.Get("tgid").(string))
	return "hello world"
}
