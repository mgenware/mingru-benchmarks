package main

import (
	"fmt"

	"github.com/mgenware/mingru-benchmarks/gormExample"
)

func main() {
	conn := gormExample.GetConn()
	rows := gormExample.SelectRows(conn)
	for _, r := range rows {
		fmt.Println(r)
	}
}
