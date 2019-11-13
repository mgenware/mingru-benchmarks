package main

import (
	"testing"

	"github.com/mgenware/mingru-benchmarks/mingruExample"
	"github.com/mgenware/mingru-benchmarks/mingruExample/da"
)

const SelectLimit = 10000

func BenchmarkMingruSelectRows(b *testing.B) {
	mrConn := mingruExample.GetConn()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, err := da.Employee.SelectAll(mrConn, 1, SelectLimit)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMingruSelectRowsWithRelationship(b *testing.B) {
	mrConn := mingruExample.GetConn()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, err := da.EmploymentHistoryItem.SelectAllHistory(mrConn, 1, SelectLimit)
		if err != nil {
			panic(err)
		}
	}
}
