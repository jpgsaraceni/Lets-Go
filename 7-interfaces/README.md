# Interfaces

Similar to structs, but interfaces contain method signatures instead of fields.

```go
type InterfaceName interface {
    SomeMethod()
}
```

Structs implement interfaces implicitly in Go. A struct satisfies an interface if it implements all its methods.

```go
type SomeInterface interface {
    SomeMethod()
}

type SomeStruct struct {
    SomeField int
}

func (s SomeStruct) SomeMethod() { // means SomeStruct satisfies SomeInterface
// because it implements all SomeInterface methods
    // do something
}

func DoSomething(someInstance SomeInterface) {
    someInstance.SomeMethod()
}

func main() {
    someVar := SomeStruct {SomeField : 1}
    DoSomething(someVar)
}
```
