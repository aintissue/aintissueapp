package app

import (
	"encoding/json"

	macaron "gopkg.in/macaron.v1"
)

// Pretty print object or variable
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func getIp(ctx *macaron.Context) string {
	ip := ""

	if len(ctx.Req.Header["X-Forwarded-For"]) > 0 {
		ip = ctx.Req.Header["X-Forwarded-For"][0]
	} else if len(ctx.Req.Header["X-Real-Ip"]) > 0 {
		ip = ctx.Req.Header["X-Real-Ip"][0]
	} else {
		ip = ctx.RemoteAddr()
	}

	return ip
}
