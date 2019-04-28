package slices

// FUNCTION "Contains" - check whether the provided slice contains the provided item

func ContainsByte(slice []byte, n byte) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func ContainsBool(slice []bool, n bool) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func ContainsInt(slice []int, n int) bool { // Checks array of ints for int
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func ContainsFloat32(slice []float32, n float32) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func ContainsFloat64(slice []float64, n float64) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func ContainsString(slice []string, n string) bool { // Checks array of strings for string
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

// FUNCTION "Count" - returns the number of instances of given item

func CountByte(slice []byte, n byte) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func CountBool(slice []bool, n bool) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func CountInt(slice []int, n int) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func CountFloat32(slice []float32, n float32) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func CountFloat64(slice []float64, n float64) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func CountString(slice []string, n string) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

// FUNCTION "Delete" - delete an item with index within given slice and return the new slice

func DeleteBool(slice []bool, i int) []bool { 
	return append(slice[:i], slice[i+1:]...)
}

func DeleteInt(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func DeleteFloat32(slice []float32, i int) []float32 {
	return append(slice[:i], slice[i+1:]...)
}

func DeleteFloat64(slice []float64, i int) []float64 {
	return append(slice[:i], slice[i+1:]...)
}

func DeleteString(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

// FUNCTION "Index" - returns index of item within given slice

func IndexByte(slice []byte, n byte) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func IndexBool(slice []bool, n bool) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func IndexInt(slice []int, n int) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func IndexFloat32(slice []float32, n float32) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func IndexFloat64(slice []float64, n float64) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func IndexString(slice []string, n string) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

// FUNCTION "LastIndex" - Returns the last index of item within given slice

func LastIndexByte(slice []byte, n byte) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func LastIndexBool(slice []bool, n bool) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func LastIndexInt(slice []int, n int) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func LastIndexFloat32(slice []float32, n float32) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func LastIndexFloat64(slice []float64, n float64) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func LastIndexString(slice []string, n string) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

// FUNCTION "Replace" - Replaces "old" items with "new" items and returns the new slice

func ReplaceByte(slice []byte, old, new byte) []byte {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func ReplaceBool(slice []bool, old, new bool) []bool {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func ReplaceInt(slice []int, old, new int) []int {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func ReplaceFloat32(slice []float32, old, new float32) []float32 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func ReplaceFloat64(slice []float64, old, new float64) []float64 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func ReplaceString(slice []string, old, new string) []string {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

// FUNCTION "ReplaceAll" - replaces all occurences of "old" with "new" and returns the new slice

func ReplaceAllByte(slice []byte, old, new byte) []byte {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
}

func ReplaceAllBool(slice []bool, old, new bool) []bool {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
}

func ReplaceAllInt(slice []int, old, new int) []int {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
}

func ReplaceAllFloat32(slice []float32, old, new float32) []float32 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
}

func ReplaceAllFloat64(slice []float64, old, new float64) []float64 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
}

func ReplaceAllString(slice []string, old, new string) []string {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
}