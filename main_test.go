package main

import (
	"testing"

	"github.com/mgenware/mingru-benchmarks/gormExample"
	"github.com/mgenware/mingru-benchmarks/mingruExample"
	"github.com/mgenware/mingru-benchmarks/mingruExample/da"
)

const SelectLimit = 10000

func BenchmarkGormSelectRows(b *testing.B) {
	gormConn := gormExample.GetConn()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		gormExample.SelectEmployees(gormConn, SelectLimit)
	}
}

func BenchmarkMingruSelectRows(b *testing.B) {
	mrConn := mingruExample.GetConn()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, err := da.Employee.SelectAll(mrConn, SelectLimit, 0, SelectLimit)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGormSelectRowsWithRelationship(b *testing.B) {
	gormConn := gormExample.GetConn()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		gormExample.SelectTitles(gormConn, SelectLimit)
	}
}

func BenchmarkMingruSelectRowsWithRelationship(b *testing.B) {
	mrConn := mingruExample.GetConn()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, err := da.Title.SelectAll(mrConn, SelectLimit, 0, SelectLimit)
		if err != nil {
			panic(err)
		}
	}
}
