# mingru-benchmarks
Benchmarks of mingru

## Run benchmarks
```sh
go test -bench=.
```

## Results on my machine
Go 1.11.5, MariaDB 10.3.9, macOS 10.13, 

```
goos: darwin
goarch: amd64
pkg: github.com/mgenware/mingru-benchmarks
BenchmarkGormSelect100Rows-4                                2000            844283 ns/op
BenchmarkMingruSelect100Rows-4                              3000            465910 ns/op
BenchmarkGormSelect100RowsWithRelationship-4                 500           2306596 ns/op
BenchmarkMingruSelect100RowsWithRelationship-4              2000            663603 ns/op
PASS
ok      github.com/mgenware/mingru-benchmarks   6.476s
```
