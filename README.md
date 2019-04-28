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

ContainsByte

ContainsBool

ContainsInt

ContainsFloat32

ContainsFloat64

ContainsString



CountByte

CountBool

CountInt

CountFloat32

CountFloat64

CountString


DeleteBool

DeleteInt

DeleteFloat32

DeleteFloat64

DeleteString


IndexByte

IndexBool

IndexInt

IndexFloat32

IndexFloat64

IndexString


LastIndexByte

LastIndexBool

LastIndexInt

LastIndexFloat32

LastIndexFloat64

LastIndexString


ReplaceByte

ReplaceBool

ReplaceInt

ReplaceFloat32

ReplaceFloat64

ReplaceString


ReplaceAllByte

ReplaceAllBool

ReplaceAllInt

ReplaceAllFloat32

ReplaceAllFloat64

ReplaceAllString
