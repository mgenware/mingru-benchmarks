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

func SelectEmployees(db *gorm.DB, limit int) []Employee {
	var users []Employee
	db.Limit(limit).Find(&users)
	return users
}

func SelectTitles(db *gorm.DB, limit int) []Title {
	var titles []Title
	db.Limit(limit).Preload("Employee").Find(&titles)
	return titles
}
