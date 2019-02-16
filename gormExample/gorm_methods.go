package gormExample

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mgenware/mingru-benchmarks/defs"
)

func GetConn() *gorm.DB {
	db, err := gorm.Open(defs.DBType, defs.DBConnString)
	if err != nil {
		panic(err)
	}
	return db
}

func SelectRows(db *gorm.DB, limit int) []User {
	var users []User
	db.Limit(limit).Find(&users)
	return users
}
