package services

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

func DB() *gorm.DB {
	return db
}

func OpenDB(dsn string, config *gorm.Config, maxIdleConns, maxOpenConns int, models ...interface{}) (err error) {
	if config == nil {
		config = &gorm.Config{}
	}

	if db, err = gorm.Open(mysql.Open(dsn), config); err != nil {
		logrus.Errorf("open database failed: %v", err)
		return
	}

	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
	} else {
		logrus.Error(err)
	}

	if err = db.AutoMigrate(models...); err != nil {
		logrus.Errorf("auto migrate failed: %v", err)
	}
	return
}

func SqlNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  len(value) > 0,
	}
}
