# mingru-benchmarks
Benchmarks of mingru

## Run benchmarks
```sh
go test -bench=. -benchmem
```

## Results on my machine
Go 1.11.5, MariaDB 10.3.9, macOS 10.13, 

```
goos: darwin
goarch: amd64
pkg: github.com/mgenware/mingru-benchmarks
BenchmarkGormSelect100Rows-4                                2000            846327 ns/op          183781 B/op       3354 allocs/op
BenchmarkMingruSelect100Rows-4                              3000            508827 ns/op           11394 B/op        523 allocs/op
BenchmarkGormSelect100RowsWithRelationship-4                 500           2521229 ns/op          424235 B/op       8144 allocs/op
BenchmarkMingruSelect100RowsWithRelationship-4              2000            733188 ns/op           47652 B/op       1025 allocs/op
```
