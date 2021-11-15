# Practicing the composite data types in Go

Learned about arrays, slices, maps and structs.

## Array

```go
var arrayOfThreeStrings = [3]string
arrayOfThreeStrings = [3]string{"Index 0", "Index 1", "Index 2"}

var arrayLiteral = [2]bool{true, true}
anotherArrayLiteral := [2]int{1, 2}
```

In Go, the size of the array defines its type. A `[5]string` is of one type and a `[6]string` is of another.

## Slice

Similar to array, except slice does not have to maintain a fixed size.

```go
var emptySliceOutOfRange = []int // created with size 0, indexes are out of range
var emptySlice = make([]int, 3) // {0, 0, 0}
emptySlice[0] = 2 // {2, 0, 0}
var sliceLiteral = []int{1, 2, 3}
anotherSliceLiteral := []int{1, 2}
```

The built-in `append()` adds elements to a slice

```go
var sliceVariable = []int{1, 2}
sliceVariable = append(sliceVariable, 3) // {1, 2, 3}
```

To append another slice use `...`

```go
var firstSlice = []string{"one", "two"}
var secondSlice = []string{"three"}
firstSlice = append(firstSlice, secondSlice...)
```

To access subsets of slices, use `[START_INDEX:END_INDEX]`, where `END_INDEX` is excludent.

```go
var set = []int{1, 2, 3, 4}
var subset1 = set[0:4] // {1, 2, 3, 4}
var subset2 = set[1:len(set)-1] // {2, 3}
```

## Map

Key/Value pairs.

```go
var someMap = make(map[int]string)
someMap[1] = "one"
someMap[2] = "two"
// {1: "one", 2: "two"}

var mapLiteral = map[string]string{
    "first key" : "some value",
    "second key": "some value",
}

anotherMapLiteral := map[int]bool{
    1 : true
    25: false
    3 : false
    2 : false
}
```

A map value can be accessed using `MAP_NAME[TARGET_KEY]]` syntax.
A map element can be deleted using `delete(MAP, KEY)`.
You can check if a value exists in a map and access its value:

```go
someMap := map[string]int{
    "some key": "some value",
    "some other key": "some other value",
}
newIdentifier, existsBool := someMap["some key"] // newIdentifier == 1, existsBool == true
anotherIdentifier, anotherExists := someMap["inexistent key"] 
// anotherIdentifier == 0, anotherExists == false
```

## Struct

A list of unique fields and their types. 

```go
type NameOfYourStruct struct {
    SomeField    string
    AnotherField string
}

instanceOfYourStruct := NameOfYourStruct{SomeField: "value", AnotherField: 10}
// instanceOfYourStruct.SomeField == "value"
```
