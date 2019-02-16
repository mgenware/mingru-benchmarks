package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/mgenware/mingru-benchmarks/gormExample"
	"github.com/mgenware/mingru-benchmarks/mingruExample"
	"github.com/mgenware/mingru-benchmarks/mingruExample/da"
)

const SelectLimit = 100

func TestMain(m *testing.M) {
	db := gormExample.GetConn()
	defer db.Close()

	for _, model := range []interface{}{
		gormExample.User{}, gormExample.Post{},
	} {
		if err := db.AutoMigrate(model).Error; err != nil {
			panic(err)
		}
	}

	// Adding users
	for i := 0; i < 100; i++ {
		user := gormExample.User{Name: fmt.Sprintf(RandStringRunes(10)), Age: i}
		db.Create(&user)

		// Adding 2 posts per user
		for j := 0; j < 2; j++ {
			post := gormExample.Post{Title: RandStringRunes(20), Content: RandStringRunes(200), User: user}
			db.Create(&post)
		}
	}

	os.Exit(m.Run())
}

func BenchmarkGormSelect100Rows(b *testing.B) {
	gormConn := gormExample.GetConn()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		gormExample.SelectRows(gormConn, SelectLimit)
	}
}

func BenchmarkMingruSelect100Rows(b *testing.B) {
	mrConn := mingruExample.GetConn()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := da.Users.SelectUsers(mrConn, SelectLimit, 0)
		if err != nil {
			panic(err)
		}
	}
}
