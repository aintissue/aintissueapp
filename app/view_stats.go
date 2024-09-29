package app

import (
	"log"

	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewStats(sess session.Store, ctx *macaron.Context) {
	sr := StatsResponse{}

	var users []*User
	db.Find(&users).Count(&sr.Users)
	db.Find(&Chat{}).Count(&sr.Projects)

	for _, u := range users {
		sr.Messages += int64(u.MsgCount)
	}

	sr.Bots = 2

	log.Println(sess.ID())

	ctx.JSON(200, sr)
}

type StatsResponse struct {
	Users    int64 `json:"users"`
	Projects int64 `json:"projects"`
	Messages int64 `json:"messages"`
	Bots     int64 `json:"bots"`
}
