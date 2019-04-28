package common

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

const (
	charset   = "utf8mb4"
	collation = "utf8mb4_bin"
)

// InitDB : Opening a database and save the reference to `Database` struct.
func InitDB() *gorm.DB {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=%s&collation=%s&parseTime=True&loc=UTC",
		user,
		password,
		dbName,
		charset,
		collation,
	))
	if err != nil {
		fmt.Println("db err: ", err)
	}

	db.DB().SetMaxIdleConns(10)
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin")
	db.SingularTable(true)
	//db.LogMode(true)

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
