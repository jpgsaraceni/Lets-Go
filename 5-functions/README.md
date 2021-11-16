# Functions

Use the keyword `func` to declare a function. If a function name starts with a capital letter, it is visible outside the package.

```go
func DoSomething(number int, anotherNumber int) bool {
    // do something with params, return something or not
    // this function must return a boolean.
}
```

To receive an argument by refference (through a pointer) use * before the parameter type and & before the argument.

```go
var variable int = 1

passByReference(&variable)

func passByReference(variable *int){
    // do something to variable
}
```

## Error Handling

One of the ways to handle errors in a function is to pass a second return type (`error`) to the function, and return 0 and a `fmt.Errorf()` message and capture on function call:

```go
result, err := Divide(2, 4)
if err != nil {
    // handle error
} else {
    // use result
}

func Divide(a int, b int) (float32, error) {
    if b == 0{
        return 0, fmt.Errorf("Division by 0 is impossible.")
    }
    return a / b
}
```

## Variadic Functions

Use `...` before the parameter type in function declaration to receive any number of arguments as an iterable:

```go
func VariadicFunction(numbers ...int){
    for _, number := range numbers {
        // do something
    }
}
```
