# Testing in go

Go supplies a standard lib for testing (called `testing`).

Create a file with the same name of the file you want to test, with `_test` before the extension. This is a basic structure of a test function:

```go
package main

import "testing"

func TestFunctionName(t *testing.T) { // use TestXxx name format for this function 
    got := FunctionName() // run the function with its arguments if it takes any
    want := "What the function should return"
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

## Running a Test in Go

You need a `go.mod` file in the directory. To create one, execute `go mod init MODNAME` in your terminal.

Enter `go test` in your terminal to run the test.

## TDD

This course is based on Test Driven Development. In a nutshell, it follows the "red, green, refactor" principal, according to which you should make a test that will fail before you write your actual code, make it pass (even if you need to "sin"), then refactor until you have a code that passes your test and works as desired.

## Subtests

```go
t.Run("scenario description", func(t *testing.T) { // test different scenarios in the same function
    // subtest
}
```

## Refactoring Tests

When you have, for example, repeated assertions for the different scenarios in your test function, you can write a function, inside the test function to be called in each subtest function:

```go
assertCorrectMessage := func(t testing.TB, got, want string) { // testing.TB is an interface that
// *testing.T and *testing.B both satisfy. Use in helper functions. 
    t.Helper() // tells the compiler this function is a helper, so if the test fails it will 
    // output the line where the actual test happened, not here.
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

## Example

You can add an example function to your _test.go file. It will be executed just like other test functions, and will be added to godoc.

```go
func ExampleFunctionName() {
    result := FunctionName("arguments")
    fmt.Println(result)
}
```

## Benchmarks

Benchmarks are also written similarly to tests. They check how long a function takes to run the b.N times and assess if that time is reasonable.

```go
func BenchmarkFunctionName(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FunctionName("arguments")
    }
}
```
