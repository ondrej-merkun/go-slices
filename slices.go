package slices

// FUNCTION "Contains" - check whether the provided slice contains the provided item

func ContainsByte(slc []byte, n byte) bool {
	for _, item := range slc {
		if item == n {
			return true
		}
	}
	return false
}

func ContainsBool(slc []bool, n bool) bool {
	for _, item := range slc {
		if item == n {
			return true
		}
	}
	return false
}

func ContainsInt(slc []int, n int) bool { // Checks array of ints for int
	for _, item := range slc {
		if item == n {
			return true
		}
	}
	return false
}

func ContainsFloat32(slc []float32, n float32) bool {
	for _, item := range slc {
		if item == n {
			return true
		}
	}
	return false
}

func ContainsFloat64(slc []float64, n float64) bool {
	for _, item := range slc {
		if item == n {
			return true
		}
	}
	return false
}

func ContainsString(slc []string, s string) bool { // Checks array of strings for string
	for _, item := range slc {
		if item == s {
			return true
		}
	}
	return false
}

// FUNCTION "Count" - returns the number of instances of given item

func CountByte(slc []byte, n byte) int {
	instances := 0
	for _, item := range slc {
		if item == n {
			instances++
		}
	}
	return instances
}

func CountBool(slc []bool, n bool) int {
	instances := 0
	for _, item := range slc {
		if item == n {
			instances++
		}
	}
	return instances
}

func CountInt(slc []int, n int) int {
	instances := 0
	for _, item := range slc {
		if item == n {
			instances++
		}
	}
	return instances
}

func CountFloat32(slc []float32, n float32) int {
	instances := 0
	for _, item := range slc {
		if item == n {
			instances++
		}
	}
	return instances
}

func CountFloat64(slc []float64, n float64) int {
	instances := 0
	for _, item := range slc {
		if item == n {
			instances++
		}
	}
	return instances
}

func CountString(slc []string, n string) int {
	instances := 0
	for _, item := range slc {
		if item == n {
			instances++
		}
	}
	return instances
}

// FUNCTION "Delete" - delete an item with index within given slice and return the new slice

func DeleteBool(slc []bool, i int) []bool { 
	return append(slc[:i], slc[i+1:]...)
}

func DeleteInt(slc []int, i int) []int {
	return append(slc[:i], slc[i+1:]...)
}

func DeleteFloat32(slc []float32, i int) []float32 {
	return append(slc[:i], slc[i+1:]...)
}

func DeleteFloat64(slc []float64, i int) []float64 {
	return append(slc[:i], slc[i+1:]...)
}

func DeleteString(slc []string, i int) []string {
	return append(slc[:i], slc[i+1:]...)
}

// FUNCTION "Index" - returns index of item within given slice

func IndexByte(slc []byte, n byte) int {
	for index, item := range slc {
		if item == n {
			return index
		}
	}
	return -1
}

func IndexBool(slc []bool, n bool) int {
	for index, item := range slc {
		if item == n {
			return index
		}
	}
	return -1
}

func IndexInt(slc []int, n int) int {
	for index, item := range slc {
		if item == n {
			return index
		}
	}
	return -1
}

func IndexFloat32(slc []float32, n float32) int {
	for index, item := range slc {
		if item == n {
			return index
		}
	}
	return -1
}

func IndexFloat64(slc []float64, n float64) int {
	for index, item := range slc {
		if item == n {
			return index
		}
	}
	return -1
}

func IndexString(slc []string, n string) int {
	for index, item := range slc {
		if item == n {
			return index
		}
	}
	return -1
}

// FUNCTION "LastIndex" - Returns the last index of item within given slice

func LastIndexByte(slc []byte, n byte) int {
	lastIndex := -1
	for index, item := range slc {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func LastIndexBool(slc []bool, n bool) int {
	lastIndex := -1
	for index, item := range slc {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func LastIndexInt(slc []int, n int) int {
	lastIndex := -1
	for index, item := range slc {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func LastIndexFloat32(slc []float32, n float32) int {
	lastIndex := -1
	for index, item := range slc {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func LastIndexFloat64(slc []float64, n float64) int {
	lastIndex := -1
	for index, item := range slc {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func LastIndexString(slc []string, n string) int {
	lastIndex := -1
	for index, item := range slc {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

// FUNCTION "Replace" - Replaces "old" items with "new" items and returns the new slice

func ReplaceByte(slc []byte, old, new byte) []byte {
	for index, item := range slc {
		if item == old {
			slc[index] = new
			return slc
		}
	}
	return slc
}

func ReplaceBool(slc []bool, old, new bool) []bool {
	for index, item := range slc {
		if item == old {
			slc[index] = new
			return slc
		}
	}
	return slc
}

func ReplaceInt(slc []int, old, new int) []int {
	for index, item := range slc {
		if item == old {
			slc[index] = new
			return slc
		}
	}
	return slc
}

func ReplaceFloat32(slc []float32, old, new float32) []float32 {
	for index, item := range slc {
		if item == old {
			slc[index] = new
			return slc
		}
	}
	return slc
}

func ReplaceFloat64(slc []float64, old, new float64) []float64 {
	for index, item := range slc {
		if item == old {
			slc[index] = new
			return slc
		}
	}
	return slc
}

func ReplaceString(slc []string, old, new string) []string {
	for index, item := range slc {
		if item == old {
			slc[index] = new
			return slc
		}
	}
	return slc
}

// FUNCTION "ReplaceAll" - replaces all occurences of "old" with "new" and returns the new slice

func ReplaceAllByte(slc []byte, old, new byte) []byte {
	for index, item := range slc {
		if item == old {
			slc[index] = new
		}
	}
	return slc
}

func ReplaceAllBool(slc []bool, old, new bool) []bool {
	for index, item := range slc {
		if item == old {
			slc[index] = new
		}
	}
	return slc
}

func ReplaceAllInt(slc []int, old, new int) []int {
	for index, item := range slc {
		if item == old {
			slc[index] = new
		}
	}
	return slc
}

func ReplaceAllFloat32(slc []float32, old, new float32) []float32 {
	for index, item := range slc {
		if item == old {
			slc[index] = new
		}
	}
	return slc
}

func ReplaceAllFloat64(slc []float64, old, new float64) []float64 {
	for index, item := range slc {
		if item == old {
			slc[index] = new
		}
	}
	return slc
}

func ReplaceAllString(slc []string, old, new string) []string {
	for index, item := range slc {
		if item == old {
			slc[index] = new
		}
	}
	return slc
}