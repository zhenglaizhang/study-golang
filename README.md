# study-golang



## Testing

- `*_test.go` for filename, prefix test function with `Test`
- accept one parameter `t testing.T`
- same package for whitebox tests (implementation)
- add `_test` suffx to the package for blackbox tests (external api, general good practice, from consumer perspective, do not care internal implementation)
- `go help testfunc`
  - test functions
  - benchmark functions
  - example functions
- leverage go test caching mechanism
- `go test -v . -run ^(TestSubstraction)$`
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