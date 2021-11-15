# Practicing the basic data types in Go.

When you give a variable an initial value, it is called a *literal*. When you don't, it receives it's *zero value*. For numerics, `0`. For strings, `""`. For booleans, `false`.

## Numeric Types

Go has three types of numeric values: integer, float point and complex. I learned about `int` and `float`. You can specify how many bits your variable can support using at the end of the type (`int8`, `float32`, ...) and if you want an absolute value put a u before the type (`uint`, `ufloat`).

## Strings

Strings can be interpreted (when declared using ") or raw (declared using `). Interpreted strings will read escape characters, that begin with \.

```go
"This\nbreaks the line."
`This \n doesn't.`
```

## Booleans

Apperently no difference from other languages. `bool` variables can have `true` or `false` as their values.