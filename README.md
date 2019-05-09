# go-slices
Unlike many other programming lanugages, Go doesn't provide helper functions for slices in it's core. I felt like this was quite an essential feature and there weren't any libraries out there that would do this, so I made my own. The functions are the same as the in the Go's official "strings" package, but you can use them on slices.


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

### Contains
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

ContainsInt8

ContainsInt16

ContainsInt32

ContainsInt64

ContainsUint

ContainsUint8

ContainsUint16

ContainsUint32

ContainsUint64

ContainsUintptr

ContainsFloat32

ContainsFloat64

ContainsString

ContainsComplex64

ContainsComplex128

---

### Count
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

CountInt8

CountInt16

CountInt32

CountInt64

CountUint

CountUint8

CountUint16

CountUint32

CountUint64

CountUintptr

CountFloat32

CountFloat64

CountString

CountComplex64

CountComplex128

---

### Delete
Deletes an item with the index (second parameter) within the slice (first parameter) and returns the new slice.
Example:
```
mySlice := []int{1, 2, 3, 4, 5}
slices.DeleteInt(mySlice, 0) // returns [2, 3, 4, 5]
```

DeleteByte

DeleteBool

DeleteInt

DeleteInt8

DeleteInt16

DeleteInt32

DeleteInt64

DeleteUint

DeleteUint8

DeleteUint16

DeleteUint32

DeleteUint64

DeleteUintptr

DeleteFloat32

DeleteFloat64

DeleteString

DeleteComplex64

DeleteComplex128

---

### Index
Returns the index of the first occurence of the item (second parameter) within the slice (first parameter).
Example:
```
mySlice := []int{4, 1, 5}
slices.IndexInt(mySlice, 5) // returns 2
```

IndexByte

IndexBool

IndexInt

IndexInt8

IndexInt16

IndexInt32

IndexInt64

IndexUint

IndexUint8

IndexUint16

IndexUint32

IndexUint64

IndexUintptr

IndexFloat32

IndexFloat64

IndexString

IndexComplex64

IndexComplex128

---

### LastIndex
Returns the index of the last occurence of the item (second parameter) within the slice (first parameter).
Example:
```
mySlice := []int{1, 2, 1, 1, 4}
slices.LastIndex(mySlice, 1) // returns 3
```

LastIndexByte

LastIndexBool

LastIndexInt

LastIndexInt8

LastIndexInt16

LastIndexInt32

LastIndexInt64

LastIndexUint

LastIndexUint8

LastIndexUint16

LastIndexUint32

LastIndexUint64

LastIndexUintptr

LastIndexFloat32

LastIndexFloat64

LastIndexString

LastIndexComplex64

LastIndexComplex128

---

### Replace
Replaces the first instance of the old item (second parameter) within the slice (first parameter) with the new item (third parameter) and returns the new slice.
Example:
```
mySlice := []int{1, 1, 3, 5, 1}
slices.ReplaceInt(mySlice, 1, 0) // returns [0, 1, 3, 5, 1]
```

ReplaceByte

ReplaceBool

ReplaceInt

ReplaceInt8

ReplaceInt16

ReplaceInt32

ReplaceInt64

ReplaceUint

ReplaceUint8

ReplaceUint16

ReplaceUint32

ReplaceUint64

ReplaceUintptr

ReplaceFloat32

ReplaceFloat64

ReplaceString

ReplaceComplex64

ReplaceComplex128

---

### ReplaceAll
Replaces every instance of the old item (second parameter) within the slice (first parameter) with the new item (third parameter) and returns the new slice.
Example:
```
mySlice := []int{1, 1, 3, 5, 1}
slices.ReplaceAllInt(mySlice, 1, 0) // returns [0, 0, 3, 5, 1]
```

ReplaceAllByte

ReplaceAllBool

ReplaceAllInt

ReplaceAllInt8

ReplaceAllInt16

ReplaceAllInt32

ReplaceAllInt64

ReplaceAllUint

ReplaceAllUint8

ReplaceAllUint16

ReplaceAllUint32

ReplaceAllUint64

ReplaceAllUintptr

ReplaceAllFloat32

ReplaceAllFloat64

ReplaceAllString

ReplaceAllComplex64

ReplaceAllComplex128
