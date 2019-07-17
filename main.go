package main

import (
	"fmt"

	"github.com/mgenware/mingru-benchmarks/gormExample"
	"github.com/mgenware/mingru-benchmarks/mingruExample"
	"github.com/mgenware/mingru-benchmarks/mingruExample/da"
)

func main() {
	limit := 10000

	fmt.Println("SelectRows - gorm")
	gormConn := gormExample.GetConn()
	users1 := gormExample.SelectEmployees(gormConn, limit)
	if len(users1) != limit {
		panic("Limit mismatch")
	}
	for _, r := range users1 {
		fmt.Println(r)
	}

	fmt.Println("SelectRows - mingru")
	mrConn := mingruExample.GetConn()
	users2, _, err := da.Employee.SelectAll(mrConn, limit, 0)
	if err != nil {
		panic(err)
	}
	if len(users2) != limit {
		panic("Limit mismatch")
	}
	for _, r := range users2 {
		fmt.Println(r)
	}

	fmt.Println("SelectRowsWithRelationship - gorm")
	titles1 := gormExample.SelectTitles(gormConn, limit)
	if len(titles1) != limit {
		panic("Limit mismatch")
	}
	for _, r := range titles1 {
		fmt.Println(r.Employee)
	}

	fmt.Println("SelectRowsWithRelationship - mingru")
	titles2, err := da.Title.SelectAll(mrConn, limit, 0)
	if err != nil {
		panic(err)
	}
	if len(titles2) != limit {
		panic("Limit mismatch")
	}
	for _, r := range titles2 {
		fmt.Println(r)
	}
}
