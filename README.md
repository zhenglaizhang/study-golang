# study-golang



## Testing

- golang.org/pkg/testing
- `*_test.go` for filename, prefix test function with `Test`
- accept one parameter `t testing.T`
- same package for whitebox tests (implementation)
- add `_test` suffx to the package for blackbox tests (external api, general good practice, from consumer perspective, do not care internal implementation)
- `go help test` & `go help testflag`
- `go help testfunc`
  - test functions
  - benchmark functions
  - example functions
- leverage go test caching mechanism
- `go test -v ./... -run ^(TestSubstraction)$`
- `go test -v ./... -run Substraction -cover`
- `go test -v net/http`
- `go test . -coverprofile cover.out` 
- `go test . -coverprofile count.out -covermode count`
- `go tool cover -func cover.out`
- `go tool cover -html cover.out`
- testing-related packages
  - testing
  - testing/quick
  - testing/iotest
  - net/http/httptest
  - testify
  - ginkgo
  - goconvey
  - ... (TODO)
- test coverage
- table driven tests
- reporting test failures
  - immediate failure (e.g. connect to db for e2e test)
    - `t.FailNow()`
    - `t.Fatal(args ...interface{})`
    - `t.Fatal(format string, args ...interface{})`
  - non-immediate failure
    - `t.Fail()`
    - `t.Error(args ...interface{})`
    - `t.Errorf(format string, args ...interface{})`
- Log & Logf
- Helper 
- Skip, Skipf, SkipNow
- Run
- Parallel

## Benchmark tests

- Prefix tests with `Benchmark`
- Accept one parameter `*testing.B`
- key testing.B memebers
  - `b.N`
  - `b.StartTimer`, e.g. exclude app startup time
  - `b.StopTimer`
  - `b.ResetTimer`
  - `b.RunParallel`
- `go test -v -bench . -benchtime 10s`
- `go test -v -bench . `, by default the benchmark time is 1s

## Profiling tests

- `go test -benchmem`
- `go test -trace {trace.out}`
- `go test -{type}profile {file}`, type listed as below
  - block
  - cover
  - cpu
  - mem
  - mutex