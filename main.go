package main

import (
	"fmt"

	"github.com/mgenware/mingru-benchmarks/gormExample"
	"github.com/mgenware/mingru-benchmarks/mingruExample"
	"github.com/mgenware/mingru-benchmarks/mingruExample/da"
)

func main() {
	limit := 100

	fmt.Println("SelectRows - gorm")
	gormConn := gormExample.GetConn()
	rows1 := gormExample.SelectRows(gormConn, limit)
	for _, r := range rows1 {
		fmt.Println(r)
	}

	fmt.Println("SelectRows - mingru")
	mrConn := mingruExample.GetConn()
	rows2, err := da.User.SelectUsers(mrConn, limit, 0)
	if err != nil {
		panic(err)
	}
	for _, r := range rows2 {
		fmt.Println(r)
	}
}
