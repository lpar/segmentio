Benchmarking https://github.com/segmentio/encoding, stdlib JSON, and https://github.com/json-iterator/go

The testdata/code.json.gz file is the benchmarking data from the segmentio/encoding repo

```
> go test -bench='BenchmarkUnmarshal.*' -benchmem -benchtime 3s -cpu 1
goos: linux
goarch: amd64
pkg: test/segmentio
cpu: Intel(R) Xeon(R) E-2276M  CPU @ 2.80GHz
BenchmarkUnmarshalStdlib    	     100	  30539353 ns/op	 7927619 B/op	  271274 allocs/op
BenchmarkUnmarshalJSONiter  	     139	  25702355 ns/op	 9244709 B/op	  346669 allocs/op
BenchmarkUnmarshalSegmentIO 	      48	  68100666 ns/op	 9283038 B/op	  243677 allocs/op
PASS
ok  	test/segmentio	12.621s
```

In case it's because of compiler/stdlib changes:

```
> go1.16.2 test -bench='BenchmarkUnmarshal.*' -benchmem -benchtime 3s -cpu 1
goos: linux
goarch: amd64
pkg: test/segmentio
cpu: Intel(R) Xeon(R) E-2276M  CPU @ 2.80GHz
BenchmarkUnmarshalStdlib    	     100	  32820881 ns/op	 7918436 B/op	  271273 allocs/op
BenchmarkUnmarshalJSONiter  	     126	  28387456 ns/op	 9235504 B/op	  346668 allocs/op
BenchmarkUnmarshalSegmentIO 	      43	  77761908 ns/op	 8975691 B/op	  230871 allocs/op
PASS
ok  	test/segmentio	13.178s
```
