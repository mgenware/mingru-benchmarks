package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestMain(m *testing.M) {
	db, err := gorm.Open("mysql", "root:123456@/mingru_benchmarks?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for _, model := range []interface{}{
		User{},
	} {
		if err := db.AutoMigrate(model).Error; err != nil {
			panic(err)
		}
	}

	// Adding users
	for i := 0; i < 100; i++ {
		user := User{Name: fmt.Sprintf("u-%v", i), Age: i}
		db.NewRecord(user)
		db.Create(&user)
	}

	os.Exit(m.Run())
}
