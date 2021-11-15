# Conditionals

## if

```go
if true { // parenthesis for the condition are optional and curly brackets are required.
    // whatever
} else if true {
    // whatever
} else {
    // whatever
}
```

It is possible to use a simple statement (scoped in if-else block) before the condition, that will be executed before evaluating the condition:

```go
if someValue := 10; true {
    // whatever
}
```

It is more idiomatic to use `return` instead of an `else` block.

## switch

```go
switch { // may use an expression after switch.
// case expressions must match switch expression (defaults to true) to be executed. 
case true: // do something
case false: // do nothing
default: // if no case matches
}
```