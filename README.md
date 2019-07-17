# mingru-benchmarks

Benchmarks of mingru

## Run benchmarks

- Install the [employees sample database](https://github.com/datacharmer/test_db)
- Make sure connection to local MySQL instance with `root:123456@/employees?parseTime=True` succeeds

To run the benchmark:

```sh
# Navigate to the root dir of this repo
go test -bench=. -benchmem
```

## Results on my machine

Go 1.12.5, MariaDB 10.3.9, macOS 10.13.

Selecting 10,000 rows:

```
goos: darwin
goarch: amd64
pkg: github.com/mgenware/mingru-benchmarks
BenchmarkGormSelectRows-4                             20          91268670 ns/op    31296982 B/op     570105 allocs/op
BenchmarkMingruSelectRows-4                          100          13304216 ns/op     3130794 B/op     100027 allocs/op
BenchmarkGormSelectRowsWithRelationship-4              5         206818964 ns/op    65786750 B/op    1044079 allocs/op
BenchmarkMingruSelectRowsWithRelationship-4           50          25187953 ns/op     3251650 B/op     110027 allocs/op
```
