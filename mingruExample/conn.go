package mingruExample

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mgenware/mingru-benchmarks/defs"
)

func GetConn() *sql.DB {
	db, err := sql.Open(defs.DBType, defs.DBConnString)
	if err != nil {
		panic(err)
	}
	return db
}
