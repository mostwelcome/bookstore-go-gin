package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
		return
	}

	DB = database
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}

// CloseDatabase Close database connection
func CloseDatabase() {
	database, error := DB.DB()

	if error != nil {
		panic("Got database error")
	}

	//close
	database.Close()
}

//// ClearTable clear table after test
//func ClearTable() {
//	DB.Exec("DELETE FROM books")
//	DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
//}

// InsertTestBook insert test book
func InsertTestBook() (Book, error) {
	b := Book{
		Author: "test",
		Title:  "test",
	}

	if err := DB.Create(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}

func ConnectTestDatabase() {
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
		return
	}

	DB = database
}
