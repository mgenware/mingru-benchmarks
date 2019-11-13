package main

import (
	"fmt"

	"github.com/mgenware/mingru-benchmarks/mingruExample"
	"github.com/mgenware/mingru-benchmarks/mingruExample/da"
)

func main() {
	limit := 10000

	fmt.Println("SelectRows")
	mrConn := mingruExample.GetConn()
	users, _, err := da.Employee.SelectAll(mrConn, 1, limit)
	if err != nil {
		panic(err)
	}
	if len(users) != limit {
		panic("Limit mismatch")
	}
	for _, r := range users {
		fmt.Println(r)
	}

	fmt.Println("SelectRowsWithRelationship")
	history, _, err := da.EmploymentHistoryItem.SelectAllHistory(mrConn, 1, limit)
	if err != nil {
		panic(err)
	}
	if len(history) != limit {
		panic("Limit mismatch")
	}
	for _, r := range history {
		fmt.Println(r)
	}
}
