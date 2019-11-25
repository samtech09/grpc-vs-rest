## grpc-vs-rest
It is just PoC to compare and benchmark Grpc vs Rest calls in golang.

### Benchmark
```
BenchmarkGetDetailGrpc-8                    9079            119139 ns/op           10046 B/op        185 allocs/op
BenchmarkGetDetailRest-8                  260584              6150 ns/op            3217 B/op         26 allocs/op
BenchmarkGetDetailRest2Live-8              16202             73938 ns/op            6861 B/op         75 allocs/op
```

**BenchmarkGetDetailRest** is written using `httptest` package.

**BenchmarkGetDetailRest2Live** is written using `http.Client` as we do in real-world when calling APIs.
