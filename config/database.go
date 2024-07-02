package config

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) *gorm.DB {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Log.WithFields(logrus.Fields{
			"DBURL": dsn,
			"error": err,
		}).Warn("Go App db Connection error")
	}
	return DB
}
