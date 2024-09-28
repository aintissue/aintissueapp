package app

import "gorm.io/gorm"

var conf *Config

var db *gorm.DB

// Package init function
func init() {
	conf = initConfig()

	db = initDb()
}

// Prepares the environment and runs the bot
func Run() {
	initMacaron()
}
