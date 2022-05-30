package db

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	// InitDBMain is a variable to be used to conneting the db main
	InitDBMain string

	// InitDBLogMongo is a variable to be used to conneting the db log
	InitDBLogMongo string
	// DebugMode is indicaror of database debug
	DebugMode bool

	initVars = true
)

// MainDB is used to access the main DB
func MainDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", InitDBMain)
	db.LogMode(true)
	if initVars {
		db.Raw("SET GLOBAL sql_mode = '';SET sql_mode = '';").Begin()
		initVars = false
	}

	if err == nil {
		return db, nil
	}

	return nil, errors.New("can't connect to main database")
}
