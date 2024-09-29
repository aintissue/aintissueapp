package app

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func initMacaron() {
	m := macaron.Classic()

	m.Use(macaron.Renderer())
	m.Use(cache.Cacher())
	m.Use(session.Sessioner())
	m.Use(macaron.Static("static"))
	m.Use(func(sess session.Store) {
		// log.Println(sess.ID())
	})

	m.Get("/stats.json", viewStats)
	m.Get("/login", viewLogin)
	m.Get("/login/:telegramid", viewLoginApi)
	m.Get("/login/:telegramid/:sessionid", viewloginDo)
	m.Get("/logout", viewLogout)

	m.Get("/", checkUser, viewApp)

	m.Run("127.0.0.1", Port)
}

func checkUser(sess session.Store, ctx *macaron.Context) {
	tgid := sess.Get("tgid")
	if tgid == nil {
		ctx.Redirect("/login")
		return
	}

	u := getUser(tgid.(int64))
	if u.ID == 0 {
		ctx.Redirect("/login")
		return
	}

	ctx.Data["User"] = u
}
