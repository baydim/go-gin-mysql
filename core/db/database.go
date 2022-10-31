package core

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ...

// db, err := sql.Open("mysql", "user:password@/dbname")
//
//	if err != nil {
//		panic(err)
//	}
//
// // See "Important settings" section.
// db.SetConnMaxLifetime(time.Minute * 3)
// db.SetMaxOpenConns(10)
// db.SetMaxIdleConns(10)
func GetConnection() *sql.DB {
	log.Println("Testing Connection")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/employe")
	if err != nil {
		panic(err)
	}
	return db
}
