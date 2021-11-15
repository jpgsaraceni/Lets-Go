# My first steps in Golang

Course organized by [Women Who Go Curitiba](https://womenwhogocwb.github.io/) available on [this Gitbook](https://womenwhogocwb.gitbook.io/letsgo/)

## Running the programs

I used [The Go Playground](https://play.golang.org/) to test my codes.

## What I've used so far

I'll try to summarize what I learned and practiced during the course:

### Package

Every Go file will start with the reserved word `package` followed by the name of the package that is written on your file. 

```go
package main
```

is used to compile the file as an executable.

### Imports

After naming your package, you can import packages that will be used in your code.

```go
import "fmt"
```

is a package for formatting input and output.

### func main()

Entry point of an executable.

```go
func main() {
    // code your program will run
}
```

### Basic Data Types

Numeric types (`int`, `float`, `complex`), character sequences (`string`) and booleans (`bool`).

#### Declaring variables

There are two ways to declare variables in Go: using the reserved word `var`, or the short declaration operator `:=`.

When using `var` you can assign an initial value to an identifier (if you don't it will be initialized with its *zero value*). You can also declare the data type. If you don't, it will be inferred. If the variable is declared outside the scope of a function, it *must* be declared with `var`.

```go
// basic syntax
var VARIABLE_NAME [TYPE] [= INITIAL_VALUE]

// example
var name string = "João"
var name string // initialized as ""
var name = "João" // type string is inferred
```

When using `:=` the type is always inferred and the initial value must be assigned.

```go
name := "João"
```

You can declare more than one variables at the same time. 

```go
var name, favoriteTeam string = "João", "Fluminense"
```

### Composite Data Types

Array, slice, map and struct. Check my notes [here](https://github.com/jpgsaraceni/Lets-Go/2-composite-types)
