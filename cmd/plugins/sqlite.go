package plugins

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqlite() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
