package app

import macaron "gopkg.in/macaron.v1"

func viewStats(ctx *macaron.Context) string {
	// lr := NotificationResponse{}
	// err := logTelegramService(ctx.Params("message"))
	// lr.Success = err == nil
	// ctx.JSON(200, lr)
	return "Hello World"
}
