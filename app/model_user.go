package app

import (
	"gorm.io/gorm"
)

// Struct representing User object
type User struct {
	gorm.Model
	TelegramId    int64  `gorm:"uniqueIndex"`
	TelUsername   string `gorm:"size:255"`
	TelFirstName  string `gorm:"size:255"`
	TelLastName   string `gorm:"size:255"`
	DefaultChatID uint
	DefaultChat   *Chat `gorm:"default:1"`
	MsgCount      uint64
	Email         string `gorm:"size:255"`
}

// Fetches User object by Telegram ID
func getUser(tid int64) *User {
	u := &User{}
	// log.Println(prettyPrint(u))
	if err := db.First(u, &User{TelegramId: tid}).Error; err != nil {
		loge(err)
	}

	return u
}
