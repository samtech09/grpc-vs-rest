## grpc-vs-rest
It is just PoC to compare and benchmark Grpc vs Rest calls in golang using different eccodings.

### Benchmark
```
goos: linux
goarch: amd64
pkg: github.com/samtech09/grpc-vs-rest/server
BenchmarkGetDetailGrpc-8                            9166            109297 ns/op           10049 B/op        185 allocs/op
BenchmarkGetDetailRestLiveByPostJsoniter-8         14172             84467 ns/op            9370 B/op        113 allocs/op
BenchmarkGetDetailRestLiveByPostMsgpack-8          10000            102643 ns/op           11046 B/op        142 allocs/op
BenchmarkGetDetailRestLiveByPost-8                 12664             92779 ns/op            9222 B/op        107 allocs/op
```

**BenchmarkGetDetailGrpc** is for grpc using protobuf

**BenchmarkGetDetailRestLiveByPostJsoniter** is using `json-iterator` for json encoding/decoding.

**BenchmarkGetDetailRestLiveByPostMsgpack** is using `messagepack` instead of json.

**BenchmarkGetDetailRestLiveByPost** is using standard json encoding/decoding.


Benchmark and tests are using `http.Client` instead of 'httptest' to simulate real network calls to endpoints for real-world comparison.
