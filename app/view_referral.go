package app

import (
	macaron "gopkg.in/macaron.v1"
)

func viewReferral(ctx *macaron.Context) {
	ctx.HTML(200, "ref")
}
