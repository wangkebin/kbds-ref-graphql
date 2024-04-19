package database

import (
	config "github.com/wangkebin/kbds-ref-graphql/configmodel"

	log "github.com/wangkebin/kbds-ref-graphql/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gorml "gorm.io/gorm/logger"
)

var db *gorm.DB

func Connect(l *log.Log) *gorm.DB {
	var err error
	if db == nil {
		db, err = gorm.Open(mysql.Open(config.GlobalConfig.ConnString), &gorm.Config{Logger: gorml.Default.LogMode(gorml.Error)})

		if err != nil {
			l.Error(err)
		}
	}
	return db
}
