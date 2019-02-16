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
BenchmarkGormSelect100Rows-4                2000            789445 ns/op
BenchmarkMingruSelect100Rows-4              3000            465390 ns/op
PASS
ok      github.com/mgenware/mingru-benchmarks   3.472s
```
