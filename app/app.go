package app

import (
	"log"

	"gorm.io/gorm"
)

var conf *Config

var db *gorm.DB

// Package init function
func init() {
	conf = initConfig()

	db = initDb()

	initKeyValue()
}

// Prepares the environment and runs the bot
func Run() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	initMacaron()
}
