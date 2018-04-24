package config

import (
	"log"
	"os"

	"github.com/jinzhu/configor"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

// TelegramConfig contains all Telegram-related configuration
type TelegramConfig struct {
	APIKey string
	APIUrl string
}

// FullURL Get APi url for specified action
func (t *TelegramConfig) FullURL(uri string) string {
	return t.APIUrl + t.APIKey + "/" + uri
}

// DatabaseConfig contains all DB-connection creds/options
type DatabaseConfig struct {
	sqlite.ConnectionURL
	Connection sqlbuilder.Database
}

// Connect establishes DB connection and runs an SQL-query
func (db *DatabaseConfig) Connect() {
	sess, err := sqlite.Open(db)
	if err != nil {
		log.Panicf("db.Open(): %q\n", err)
	}
	sess.SetLogging(true)
	db.Connection = sess
}

// YouTubeConfig holds configuration for YouTube API
type YouTubeConfig struct {
	DeveloperKey string
}

// Telegram is an instance of TelegramConfig
var Telegram TelegramConfig

// DB connection instance
var DB DatabaseConfig

// YT is a YouTube config holder
var YT YouTubeConfig

func init() {
	var tgConfigFile = "config/telegram_config.json"
	if _, err := os.Stat(tgConfigFile); os.IsNotExist(err) {
		log.Panicf("Missing Config: file %s was not found. \n", tgConfigFile)
	}
	if err := configor.Load(&Telegram, tgConfigFile); err != nil {
		log.Panicln(err)
	}

	var dbConfigFile = "config/database_config.json"
	if _, err := os.Stat(dbConfigFile); os.IsNotExist(err) {
		log.Panicf("Missing Config: file %s was not found. \n", dbConfigFile)
	}
	if err := configor.Load(&DB, dbConfigFile); err != nil {
		log.Panicln(err)
	}

	var ytConfigFile = "config/youtube_config.json"
	if _, err := os.Stat(ytConfigFile); os.IsNotExist(err) {
		log.Panicf("Missing Config: file %s was not found. \n", ytConfigFile)
	}
	if err := configor.Load(&YT, ytConfigFile); err != nil {
		log.Panicln(err)
	}
}
