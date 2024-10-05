package app

import (
	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

func viewProfile(sess session.Store, ctx *macaron.Context) {
	ctx.HTML(200, "profile")
}

func viewProfileSave(prof ProfileForm, f *session.Flash, ctx *macaron.Context) {
	if len(prof.Email) == 0 {
		f.Error("Email field is required.")
		ctx.Redirect("/profile")
		return
	} else {
		u := ctx.Data["User"].(*User)
		u.Email = prof.Email
		err := db.Save(u).Error
		if err != nil {
			loge(err)
		}
		f.Success("Email saved successfully.")
		ctx.Redirect("/profile")
		return
	}
}

type ProfileForm struct {
	Email string `form:"email"`
}
