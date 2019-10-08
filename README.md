# go-slices 
[![Go Report Card](https://goreportcard.com/badge/github.com/merkur0/go-slices)](https://goreportcard.com/report/github.com/merkur0/go-slices)
[![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Check%20out%20this%20library%20that%20makes%20your%20life%20easier%20when%20working%20with%20slices%20in%20Go!&url=https://github.com/merkur0/go-slices)

Unlike many other programming languages, Go doesn't provide helper functions for slices in it's core. I felt like this was quite an essential feature and there weren't any libraries out there that would do this, so I made my own. The functions are the same as the in the Go's official "strings" package, but you can use them on slices.


## How to install
Open your terminal and type:
```
go get github.com/merkur0/go-slices
```
Then, import this package in your .go file like this:
```go
import (
  "github.com/merkur0/go-slices"
)
```
And that's it!

## Usage:
```go
slices.FunctionName(parameters)
```  
  
Example:
```go
mySlice := []int{1, 2, 3}
slices.ContainsInt(mySlice, 3) // Returns "true"
```


## List of functions:

### Contains
Returns true if the slice contains the item and false if it doesn't.
Example:
```go
mySlice := []float32{1.1, 1.2, 1.3, 1.4, 1.5}
slices.ContainsFloat32(myArr, 1.2) // returns true
slices.ContainsFloat32(myArr, 2.3) // returns false
```

---

### Count
Returns the number of instances of the item within the slice.
Example:
```go
mySlice := []int{1, 1, 1, 2, 1}
slices.CountInt(mySlice, 1) // returns 4
slices.CountInt(mySlice, 3) // returns 0
```

---

### Delete
Returns a copy of the slice where the item with the index (second parameter) is deleted.
Example:
```go
mySlice := []int{1, 2, 3, 4, 5}
slices.DeleteInt(mySlice, 0) // returns [2, 3, 4, 5]
```

---

### Index
Returns the index of the first instance of the second parameter in the slice, or -1 if item is not present in the slice.
Example:
```go
mySlice := []int{4, 1, 5}
slices.IndexInt(mySlice, 5) // returns 2
```

---

### LastIndex
Returns the index of the last instance of the second parameter in the slice, or -1 if item is not present in the slice.
Example:
```go
mySlice := []int{1, 2, 1, 1, 4}
slices.LastIndexInt(mySlice, 1) // returns 3
```

---

### Map
Returns a copy of the slice with all its items modified according to the mapping function.
Example:
```go
myFunc := func(number int) int {
    return number * 2
}
mySlice := []int{1, 2, 3, 4}

slices.MapInt(mySlice, myFunc) // returns [2, 4, 6, 8]
```

---

### Max
Returns the largest element in the slice.
Example:
```go
mySlice := []int{1, 2, 3, 4}

slices.MaxInt(mySlice) // returns 4
```

---

### Min
Returns the smallest element in the slice.
Example:
```go
mySlice := []int{1, 2, 3, 4}

slices.MinInt(mySlice) // returns 1
```

---

### Pop
Returns a copy of the slice with it's last item removed.
Example:
```go
mySlice := []int{1, 2, 3, 4}
slices.PopInt(mySlice) // returns [1, 2, 3]
```

---

### Replace
Returns a copy of the slice where the first instance of the old item (second parameter) is replaced with the new item (third parameter).
Example:
```go
mySlice := []int{1, 1, 3, 5, 1}
slices.ReplaceInt(mySlice, 1, 0) // returns [0, 1, 3, 5, 1]
```

---

### ReplaceAll
Returns a copy of the slice where every instance of the old item (second parameter) is replaced with the new item (third parameter) and returns the new slice.
Example:
```go
mySlice := []int{1, 1, 3, 5, 1}
slices.ReplaceAllInt(mySlice, 1, 0) // returns [0, 0, 3, 5, 1]
```

---


### Shift
Returns a copy of the slice with it's first item removed.
Example:
```go
mySlice := []int{1, 2, 3}
slices.ShiftInt(mySlice) // returns [2, 3]
```

---
