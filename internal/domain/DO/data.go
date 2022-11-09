package DO

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Newdata() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//db.AutoMigrate(&library.Lookbook{})
	return db
}
