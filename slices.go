package slices

func ContainsByte(slice []byte, n byte) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountByte(slice []byte, n byte) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteByte(slice []byte, i int) []byte {
	return append(slice[:i], slice[i+1:]...)
}

func IndexByte(slice []byte, n byte) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexByte(slice []byte, n byte) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapByte(slice []byte, f func(byte) byte) []byte {
	var newSlice []byte
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceByte(newSlice, item, f(item))
	}
	return newSlice
}

func PopByte(slice []byte) []byte {
	return slice[:(len(slice) - 1)]
}

func ReplaceByte(slice []byte, old, new byte) []byte {
	var newSlice []byte
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllByte(slice []byte, old, new byte) []byte {
	var newSlice []byte
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftByte(slice []byte) []byte {
	return slice[1:]
}

func ContainsBool(slice []bool, n bool) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
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

func DeleteBool(slice []bool, i int) []bool {
	return append(slice[:i], slice[i+1:]...)
}

func IndexBool(slice []bool, n bool) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
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

func MapBool(slice []bool, f func(bool) bool) []bool {
	var newSlice []bool
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceBool(newSlice, item, f(item))
	}
	return newSlice
}

func PopBool(slice []bool) []bool {
	return slice[:(len(slice) - 1)]
}

func ReplaceBool(slice []bool, old, new bool) []bool {
	var newSlice []bool
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllBool(slice []bool, old, new bool) []bool {
	var newSlice []bool
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftBool(slice []bool) []bool {
	return slice[1:]
}

func ContainsInt(slice []int, n int) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
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

func DeleteInt(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func IndexInt(slice []int, n int) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
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

func MapInt(slice []int, f func(int) int) []int {
	var newSlice []int
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceInt(newSlice, item, f(item))
	}
	return newSlice
}

func PopInt(slice []int) []int {
	return slice[:(len(slice) - 1)]
}

func ReplaceInt(slice []int, old, new int) []int {
	var newSlice []int
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllInt(slice []int, old, new int) []int {
	var newSlice []int
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftInt(slice []int) []int {
	return slice[1:]
}

func MaxInt(slice []int) int {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinInt(slice []int) int {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsInt8(slice []int8, n int8) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountInt8(slice []int8, n int8) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteInt8(slice []int8, i int) []int8 {
	return append(slice[:i], slice[i+1:]...)
}

func IndexInt8(slice []int8, n int8) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexInt8(slice []int8, n int8) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapInt8(slice []int8, f func(int8) int8) []int8 {
	var newSlice []int8
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceInt8(newSlice, item, f(item))
	}
	return newSlice
}

func PopInt8(slice []int8) []int8 {
	return slice[:(len(slice) - 1)]
}

func ReplaceInt8(slice []int8, old, new int8) []int8 {
	var newSlice []int8
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllInt8(slice []int8, old, new int8) []int8 {
	var newSlice []int8
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftInt8(slice []int8) []int8 {
	return slice[1:]
}

func MaxInt8(slice []int8) int8 {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinInt8(slice []int8) int8 {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsInt16(slice []int16, n int16) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountInt16(slice []int16, n int16) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteInt16(slice []int16, i int) []int16 {
	return append(slice[:i], slice[i+1:]...)
}

func IndexInt16(slice []int16, n int16) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexInt16(slice []int16, n int16) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapInt16(slice []int16, f func(int16) int16) []int16 {
	var newSlice []int16
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceInt16(newSlice, item, f(item))
	}
	return newSlice
}

func PopInt16(slice []int16) []int16 {
	return slice[:(len(slice) - 1)]
}

func ReplaceInt16(slice []int16, old, new int16) []int16 {
	var newSlice []int16
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllInt16(slice []int16, old, new int16) []int16 {
	var newSlice []int16
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftInt16(slice []int16) []int16 {
	return slice[1:]
}

func MaxInt16(slice []int16) int16 {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinInt16(slice []int16) int16 {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsInt32(slice []int32, n int32) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountInt32(slice []int32, n int32) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteInt32(slice []int32, i int) []int32 {
	return append(slice[:i], slice[i+1:]...)
}

func IndexInt32(slice []int32, n int32) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexInt32(slice []int32, n int32) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapInt32(slice []int32, f func(int32) int32) []int32 {
	var newSlice []int32
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceInt32(newSlice, item, f(item))
	}
	return newSlice
}

func PopInt32(slice []int32) []int32 {
	return slice[:(len(slice) - 1)]
}

func ReplaceInt32(slice []int32, old, new int32) []int32 {
	var newSlice []int32
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllInt32(slice []int32, old, new int32) []int32 {
	var newSlice []int32
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftInt32(slice []int32) []int32 {
	return slice[1:]
}

func MaxInt32(slice []int32) int32 {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinInt32(slice []int32) int32 {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsInt64(slice []int64, n int64) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountInt64(slice []int64, n int64) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteInt64(slice []int64, i int) []int64 {
	return append(slice[:i], slice[i+1:]...)
}

func IndexInt64(slice []int64, n int64) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexInt64(slice []int64, n int64) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapInt64(slice []int64, f func(int64) int64) []int64 {
	var newSlice []int64
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceInt64(newSlice, item, f(item))
	}
	return newSlice
}

func PopInt64(slice []int64) []int64 {
	return slice[:(len(slice) - 1)]
}

func ReplaceInt64(slice []int64, old, new int64) []int64 {
	var newSlice []int64
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllInt64(slice []int64, old, new int64) []int64 {
	var newSlice []int64
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftInt64(slice []int64) []int64 {
	return slice[1:]
}

func MaxInt64(slice []int64) int64 {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinInt64(slice []int64) int64 {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsFloat32(slice []float32, n float32) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
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

func DeleteFloat32(slice []float32, i int) []float32 {
	return append(slice[:i], slice[i+1:]...)
}

func IndexFloat32(slice []float32, n float32) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
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

func MapFloat32(slice []float32, f func(float32) float32) []float32 {
	var newSlice []float32
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceFloat32(newSlice, item, f(item))
	}
	return newSlice
}

func PopFloat32(slice []float32) []float32 {
	return slice[:(len(slice) - 1)]
}

func ReplaceFloat32(slice []float32, old, new float32) []float32 {
	var newSlice []float32
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllFloat32(slice []float32, old, new float32) []float32 {
	var newSlice []float32
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftFloat32(slice []float32) []float32 {
	return slice[1:]
}

func MaxFloat32(slice []float32) float32 {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinFloat32(slice []float32) float32 {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsFloat64(slice []float64, n float64) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
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

func DeleteFloat64(slice []float64, i int) []float64 {
	return append(slice[:i], slice[i+1:]...)
}

func IndexFloat64(slice []float64, n float64) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
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

func MapFloat64(slice []float64, f func(float64) float64) []float64 {
	var newSlice []float64
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceFloat64(newSlice, item, f(item))
	}
	return newSlice
}

func PopFloat64(slice []float64) []float64 {
	return slice[:(len(slice) - 1)]
}

func ReplaceFloat64(slice []float64, old, new float64) []float64 {
	var newSlice []float64
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllFloat64(slice []float64, old, new float64) []float64 {
	var newSlice []float64
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftFloat64(slice []float64) []float64 {
	return slice[1:]
}

func MaxFloat64(slice []float64) float64 {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinFloat64(slice []float64) float64 {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsString(slice []string, n string) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
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

func DeleteString(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

func IndexString(slice []string, n string) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
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

func MapString(slice []string, f func(string) string) []string {
	var newSlice []string
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceString(newSlice, item, f(item))
	}
	return newSlice
}

func PopString(slice []string) []string {
	return slice[:(len(slice) - 1)]
}

func ReplaceString(slice []string, old, new string) []string {
	var newSlice []string
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllString(slice []string, old, new string) []string {
	var newSlice []string
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftString(slice []string) []string {
	return slice[1:]
}

func MaxString(slice []string) string {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinString(slice []string) string {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsUint(slice []uint, n uint) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountUint(slice []uint, n uint) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteUint(slice []uint, i int) []uint {
	return append(slice[:i], slice[i+1:]...)
}

func IndexUint(slice []uint, n uint) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexUint(slice []uint, n uint) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapUint(slice []uint, f func(uint) uint) []uint {
	var newSlice []uint
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceUint(newSlice, item, f(item))
	}
	return newSlice
}

func PopUint(slice []uint) []uint {
	return slice[:(len(slice) - 1)]
}

func ReplaceUint(slice []uint, old, new uint) []uint {
	var newSlice []uint
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllUint(slice []uint, old, new uint) []uint {
	var newSlice []uint
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftUint(slice []uint) []uint {
	return slice[1:]
}

func MaxUint(slice []uint) uint {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinUint(slice []uint) uint {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsUint8(slice []uint8, n uint8) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountUint8(slice []uint8, n uint8) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteUint8(slice []uint8, i int) []uint8 {
	return append(slice[:i], slice[i+1:]...)
}

func IndexUint8(slice []uint8, n uint8) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexUint8(slice []uint8, n uint8) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapUint8(slice []uint8, f func(uint8) uint8) []uint8 {
	var newSlice []uint8
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceUint8(newSlice, item, f(item))
	}
	return newSlice
}

func PopUint8(slice []uint8) []uint8 {
	return slice[:(len(slice) - 1)]
}

func ReplaceUint8(slice []uint8, old, new uint8) []uint8 {
	var newSlice []uint8
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllUint8(slice []uint8, old, new uint8) []uint8 {
	var newSlice []uint8
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftUint8(slice []uint8) []uint8 {
	return slice[1:]
}

func MaxUint8(slice []uint8) uint8 {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinUint8(slice []uint8) uint8 {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsUint16(slice []uint16, n uint16) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountUint16(slice []uint16, n uint16) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteUint16(slice []uint16, i int) []uint16 {
	return append(slice[:i], slice[i+1:]...)
}

func IndexUint16(slice []uint16, n uint16) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexUint16(slice []uint16, n uint16) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapUint16(slice []uint16, f func(uint16) uint16) []uint16 {
	var newSlice []uint16
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceUint16(newSlice, item, f(item))
	}
	return newSlice
}

func PopUint16(slice []uint16) []uint16 {
	return slice[:(len(slice) - 1)]
}

func ReplaceUint16(slice []uint16, old, new uint16) []uint16 {
	var newSlice []uint16
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllUint16(slice []uint16, old, new uint16) []uint16 {
	var newSlice []uint16
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftUint16(slice []uint16) []uint16 {
	return slice[1:]
}

func MaxUint16(slice []uint16) uint16 {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinUint16(slice []uint16) uint16 {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsUint32(slice []uint32, n uint32) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountUint32(slice []uint32, n uint32) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteUint32(slice []uint32, i int) []uint32 {
	return append(slice[:i], slice[i+1:]...)
}

func IndexUint32(slice []uint32, n uint32) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexUint32(slice []uint32, n uint32) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapUint32(slice []uint32, f func(uint32) uint32) []uint32 {
	var newSlice []uint32
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceUint32(newSlice, item, f(item))
	}
	return newSlice
}

func PopUint32(slice []uint32) []uint32 {
	return slice[:(len(slice) - 1)]
}

func ReplaceUint32(slice []uint32, old, new uint32) []uint32 {
	var newSlice []uint32
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllUint32(slice []uint32, old, new uint32) []uint32 {
	var newSlice []uint32
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftUint32(slice []uint32) []uint32 {
	return slice[1:]
}

func MaxUint32(slice []uint32) uint32 {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinUint32(slice []uint32) uint32 {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsUint64(slice []uint64, n uint64) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountUint64(slice []uint64, n uint64) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteUint64(slice []uint64, i int) []uint64 {
	return append(slice[:i], slice[i+1:]...)
}

func IndexUint64(slice []uint64, n uint64) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexUint64(slice []uint64, n uint64) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapUint64(slice []uint64, f func(uint64) uint64) []uint64 {
	var newSlice []uint64
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceUint64(newSlice, item, f(item))
	}
	return newSlice
}

func PopUint64(slice []uint64) []uint64 {
	return slice[:(len(slice) - 1)]
}

func ReplaceUint64(slice []uint64, old, new uint64) []uint64 {
	var newSlice []uint64
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllUint64(slice []uint64, old, new uint64) []uint64 {
	var newSlice []uint64
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftUint64(slice []uint64) []uint64 {
	return slice[1:]
}

func MaxUint64(slice []uint64) uint64 {
	max := slice[0]
	for _, item := range slice {
		if item > max {
			max = item
		}
	}
	return max
}

func MinUint64(slice []uint64) uint64 {
	min := slice[0]
	for _, item := range slice {
		if item < min {
			min = item
		}
	}
	return min
}

func ContainsUintptr(slice []uintptr, n uintptr) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountUintptr(slice []uintptr, n uintptr) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteUintptr(slice []uintptr, i int) []uintptr {
	return append(slice[:i], slice[i+1:]...)
}

func IndexUintptr(slice []uintptr, n uintptr) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexUintptr(slice []uintptr, n uintptr) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapUintptr(slice []uintptr, f func(uintptr) uintptr) []uintptr {
	var newSlice []uintptr
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceUintptr(newSlice, item, f(item))
	}
	return newSlice
}

func PopUintptr(slice []uintptr) []uintptr {
	return slice[:(len(slice) - 1)]
}

func ReplaceUintptr(slice []uintptr, old, new uintptr) []uintptr {
	var newSlice []uintptr
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllUintptr(slice []uintptr, old, new uintptr) []uintptr {
	var newSlice []uintptr
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftUintptr(slice []uintptr) []uintptr {
	return slice[1:]
}

func ContainsComplex64(slice []complex64, n complex64) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountComplex64(slice []complex64, n complex64) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteComplex64(slice []complex64, i int) []complex64 {
	return append(slice[:i], slice[i+1:]...)
}

func IndexComplex64(slice []complex64, n complex64) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexComplex64(slice []complex64, n complex64) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapComplex64(slice []complex64, f func(complex64) complex64) []complex64 {
	var newSlice []complex64
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceComplex64(newSlice, item, f(item))
	}
	return newSlice
}

func PopComplex64(slice []complex64) []complex64 {
	return slice[:(len(slice) - 1)]
}

func ReplaceComplex64(slice []complex64, old, new complex64) []complex64 {
	var newSlice []complex64
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllComplex64(slice []complex64, old, new complex64) []complex64 {
	var newSlice []complex64
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftComplex64(slice []complex64) []complex64 {
	return slice[1:]
}

func ContainsComplex128(slice []complex128, n complex128) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}

func CountComplex128(slice []complex128, n complex128) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}

func DeleteComplex128(slice []complex128, i int) []complex128 {
	return append(slice[:i], slice[i+1:]...)
}

func IndexComplex128(slice []complex128, n complex128) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}

func LastIndexComplex128(slice []complex128, n complex128) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}

func MapComplex128(slice []complex128, f func(complex128) complex128) []complex128 {
	var newSlice []complex128
	copy(newSlice, slice)
	for _, item := range newSlice {
		ReplaceComplex128(newSlice, item, f(item))
	}
	return newSlice
}

func PopComplex128(slice []complex128) []complex128 {
	return slice[:(len(slice) - 1)]
}

func ReplaceComplex128(slice []complex128, old, new complex128) []complex128 {
	var newSlice []complex128
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
			return newSlice
		}
	}
	return newSlice
}

func ReplaceAllComplex128(slice []complex128, old, new complex128) []complex128 {
	var newSlice []complex128
	copy(newSlice, slice)
	for index, item := range newSlice {
		if item == old {
			newSlice[index] = new
		}
	}
	return newSlice
}

func ShiftComplex128(slice []complex128) []complex128 {
	return slice[1:]
}
