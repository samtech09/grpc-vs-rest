## grpc-vs-rest
It is just PoC to compare and benchmark Grpc vs Rest calls in golang.

### Benchmark
```
goos: linux
goarch: amd64
pkg: github.com/samtech09/grpc-vs-rest/server
BenchmarkGetDetailGrpc-8                    8526            148237 ns/op           10048 B/op        185 allocs/op
BenchmarkGetDetailRest-8                  200913              5455 ns/op            3217 B/op         26 allocs/op
BenchmarkGetDetailRestLive-8               14852             74429 ns/op            6860 B/op         75 allocs/op
BenchmarkGetDetailRestLiveByPost-8         12901             94485 ns/op            9230 B/op        107 allocs/op
```

**BenchmarkGetDetailRest** is written using `httptest` package.

**BenchmarkGetDetailRest2Live** is written using `http.Client` as we do in real-world when calling APIs.

**BenchmarkGetDetailRestLiveByPost** is using marshaling/unmarshling request and response to JSON for real comparison, as grpc auto marshal/unmarshal on each request.
