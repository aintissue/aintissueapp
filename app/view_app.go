package app

import (
	"fmt"

	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewApp(sess session.Store, ctx *macaron.Context) {
	u := ctx.Data["User"].(*User)
	ctx.Data["Upgrade"] = 0

	ctx.Data["Chats"] = u.getChats()
	ctx.Data["Bots"] = u.getBots()
	ctx.Data["Stats"] = getStats(u)

	if u.Plan == PlanFree && (len(u.getChats()) > 0 || len(u.getBots()) > 0) {
		ctx.Data["Upgrade"] = 1
	}

	ctx.HTML(200, "home")
}

func getStats(u *User) *Stats {
	s := &Stats{}
	mrr := float64(0)

	mrr += (float64(u.countBasic()) * 2)
	mrr += (float64(u.countBusiness()) * 10)

	s.Referred = u.countReferred()
	s.Basic = u.countBasic()
	s.Business = u.countBusiness()
	s.MRR = fmt.Sprintf("%.2f", mrr)

	return s
}

type Stats struct {
	Referred int
	Basic    int
	Business int
	MRR      string
}
