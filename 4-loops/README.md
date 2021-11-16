# Loops

Go synthesizes all the loop structures in the `for` structure.

## `for` as a counter

```go
for i := 0; i < 10; i++ {
    // this will repeat 10 times
} 
```

## `for` with stopping criteria (like while)

```go
for someCondition == false {
    // this code will run until someCondition is true
}
```

## `for` for infinite loops

```go
for {
    // do something forever
}
```

## `for` to iterate over string, array, slice or map using `range`

```go
for INDEX[, ELEMENT] := range ITERABLE {
    // code that will be executed on each element
}
```
