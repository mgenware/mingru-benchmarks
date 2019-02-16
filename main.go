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
	users1 := gormExample.SelectUsers(gormConn, limit)
	for _, r := range users1 {
		fmt.Println(r)
	}

	fmt.Println("SelectRows - mingru")
	mrConn := mingruExample.GetConn()
	users2, err := da.Users.SelectUsers(mrConn, limit, 0)
	if err != nil {
		panic(err)
	}
	for _, r := range users2 {
		fmt.Println(r)
	}

	fmt.Println("SelectRowsWithRelationship - gorm")
	posts1 := gormExample.SelectPosts(gormConn, limit)
	for _, r := range posts1 {
		fmt.Println(r)
	}

	fmt.Println("SelectRowsWithRelationship - mingru")
	post2, err := da.Posts.SelectPosts(mrConn, limit, 0)
	if err != nil {
		panic(err)
	}
	for _, r := range post2 {
		fmt.Println(r)
	}
}
