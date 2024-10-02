package app

import (
	"log"

	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewProfile(sess session.Store, ctx *macaron.Context) {
	ctx.HTML(200, "profile")
}

func viewProfileSave(prof ProfileForm, ctx *macaron.Context) {
	log.Println(prof.Email)
	ctx.HTML(200, "profile")
}

type ProfileForm struct {
	Email string `form:"email"`
}
