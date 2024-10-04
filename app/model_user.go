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
	Email         string `gorm:"size:255 uniqueIndex"`
	RefCode       string `gorm:"size:255 uniqueIndex"`
	Plan          int    `gorm:"default:0"`
	// Chats         []*Chat
}

func (u *User) getChats() []*Chat {
	var chats []*Chat
	db.Find(&chats, &Chat{OwnerID: u.ID})
	return chats
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
