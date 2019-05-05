package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	charset   = "utf8mb4"
	collation = "utf8mb4_bin"
)

type DBConfig struct {
	User         string
	Password     string
	DatabaseName string
}

// InitDB : Opening a database and save the reference to `Database` struct.
func Init(conf DBConfig) *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=%s&collation=%s&parseTime=True&loc=UTC",
		conf.User,
		conf.Password,
		conf.DatabaseName,
		charset,
		collation,
	))
	if err != nil {
		fmt.Println("db err: ", err)
	}

	db.DB().SetMaxIdleConns(10)
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin")
	db.SingularTable(true)
	db.LogMode(true)

	return db
}
