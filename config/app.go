package config

import (
	"github.com/jinzhu/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "postgres://grealyve:123456@localhost:5432/blogdb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
