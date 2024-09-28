package app

import (
	"github.com/go-macaron/cache"
	macaron "gopkg.in/macaron.v1"
)

func initMacaron() {
	m := macaron.Classic()

	m.Use(macaron.Renderer())
	m.Use(cache.Cacher())

	// m.Get("/alpha-sent/:addr", viewAlphaSent)

	m.Run("127.0.0.1", Port)
}
