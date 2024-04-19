package database

import (
	"github.com/wangkebin/kbds-ref-graphql/graph/model"
	log "github.com/wangkebin/kbds-ref-graphql/logger"
	"gorm.io/gorm"
)

func SearchFiles(name *string, db *gorm.DB, l *log.Log) ([]*model.FileInfo, error) {
	var files []*model.FileInfo

	var res = db.Table("file_info").Where("name like ", name).Find(&files)
	if res.Error != nil {
		l.Error(res.Error)
		return nil, res.Error
	}
	return files, nil
}
