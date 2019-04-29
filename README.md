# go-slices
This package implements functions similar to the ones from the golang's official "strings" package, but for slices!


## How to install
Open your terminal and type:
```
go get github.com/merkur0/go-slices
```
Then, import this package with this code:
```
import (
  "github.com/merkur0/go-slices"
)
```
And that's it!

## Usage:
```
slices.FunctionName(parameters)
```  
  
Example:
```
mySlice := []int{1, 2, 3}
slices.ContainsInt(mySlice, 3) // Returns "true"
```


## List of functions:

#### Contains
Returns true if the slice (first parameter) contains the item (second parameter). Otherwise returns false.
Example:
```
mySlice := []float32{1.1, 1.2, 1.3, 1.4, 1.5}
slices.ContainsFloat32(myArr, 1.2) // returns true
slices.ContainsFloat32(myArr, 2.3) // returns false
```

ContainsByte

ContainsBool

ContainsInt

ContainsFloat32

ContainsFloat64

ContainsString

---

#### Count
Returns the number of instances of the item (second parameter) within the slice (first parameter).
Example:
```
mySlice := []int{1, 1, 1, 2, 1}
slices.CountInt(mySlice, 1) // returns 4
slices.CountInt(mySlice, 3) // returns 0
```

CountByte

CountBool

CountInt

CountFloat32

CountFloat64

CountString

---

#### Delete
Deletes an item with the index (second parameter) within the slice (first parameter) and returns the new slice.
Example:
```
mySlice := []int{1, 2, 3, 4, 5}
slices.DeleteInt(mySlice, 0) // returns [2, 3, 4, 5]
```

DeleteBool

DeleteInt

DeleteFloat32

DeleteFloat64

DeleteString

---

#### Index
Returns the index of the first occurence of the item (second parameter) within the slice (first parameter).
Example:
```
mySlice := []int{4, 1, 5}
slices.IndexInt(mySlice, 5) // returns 2
```

IndexByte

IndexBool

IndexInt

IndexFloat32

IndexFloat64

IndexString

---

#### LastIndex
Returns the index of the last occurence of the item (second parameter) within the slice (first parameter).
Example:
```
mySlice := []int{1, 2, 1, 1, 4}
slices.LastIndex(mySlice, 1) // returns 3
```

LastIndexByte

LastIndexBool

LastIndexInt

LastIndexFloat32

LastIndexFloat64

LastIndexString

---

#### Replace
Replaces the first instance of the old item (second parameter) within the slice (first parameter) with the new item (third parameter) and returns the new slice.
Example:
```
mySlice := []int{1, 1, 3, 5, 1}
slices.ReplaceInt(mySlice, 1, 0) // returns [0, 1, 3, 5, 1]
```

ReplaceByte

ReplaceBool

ReplaceInt

ReplaceFloat32

ReplaceFloat64

ReplaceString

---

#### ReplaceAll
Replaces every instance of the old item (second parameter) within the slice (first parameter) with the new item (third parameter) and returns the new slice.
Example:
```
mySlice := []int{1, 1, 3, 5, 1}
slices.ReplaceAllInt(mySlice, 1, 0) // returns [0, 0, 3, 5, 1]
```

ReplaceAllByte

ReplaceAllBool

ReplaceAllInt

ReplaceAllFloat32

ReplaceAllFloat64

ReplaceAllString
