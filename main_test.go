package main

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestMain(m *testing.M) {
	db, err := gorm.Open("mysql", "root:123456@/mingru_benchmarks?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for _, model := range []interface{}{
		User{}, Post{},
	} {
		if err := db.AutoMigrate(model).Error; err != nil {
			panic(err)
		}
	}

	// Adding users
	for i := 0; i < 100; i++ {
		user := User{Name: fmt.Sprintf(RandStringRunes(10)), Age: i}
		db.Create(&user)

		// Adding 2 posts per user
		for j := 0; j < 2; j++ {
			post := Post{Title: RandStringRunes(20), Content: RandStringRunes(200), User: user}
			db.Create(&post)
		}
	}

	os.Exit(m.Run())
}
