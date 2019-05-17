# mingru-benchmarks

Benchmarks of mingru

## Run benchmarks

- Install the [employees sample database](https://github.com/datacharmer/test_db)
- Make sure connection to local MySQL instance with `root:123456@/employees?parseTime=True` succeeds

To run the benchmark:

```sh
go test -bench=. -benchmem
```

## Results on my machine

Go 1.12.5, MariaDB 10.3.9, macOS 10.13.

Selecting 10,000 rows:

```
goos: darwin
goarch: amd64
pkg: github.com/mgenware/mingru-benchmarks
BenchmarkGormSelectRows-8                             20          75914389 ns/op        31297620 B/op     570107 allocs/op
BenchmarkMingruSelectRows-8                          100          10937288 ns/op         3130817 B/op     100027 allocs/op
BenchmarkGormSelectRowsWithRelationship-8             10         180885123 ns/op        65757827 B/op    1044081 allocs/op
BenchmarkMingruSelectRowsWithRelationship-8          100          15630878 ns/op         3251641 B/op     110027 allocs/op
```
