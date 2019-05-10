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

func ReplaceByte(slice []byte, old, new byte) []byte {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapByte(slice []byte, f func(byte) byte) []byte {
	for _, item := range slice {
		ReplaceByte(slice, item, f(item))
	}
	return slice
}

func ReplaceAllByte(slice []byte, old, new byte) []byte {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
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

func ReplaceBool(slice []bool, old, new bool) []bool {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapBool(slice []bool, f func(bool) bool) []bool {
	for _, item := range slice {
		ReplaceBool(slice, item, f(item))
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

func ReplaceInt(slice []int, old, new int) []int {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapInt(slice []int, f func(int) int) []int {
	for _, item := range slice {
		ReplaceInt(slice, item, f(item))
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

func ReplaceInt8(slice []int8, old, new int8) []int8 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapInt8(slice []int8, f func(int8) int8) []int8 {
	for _, item := range slice {
		ReplaceInt8(slice, item, f(item))
	}
	return slice
}

func ReplaceAllInt8(slice []int8, old, new int8) []int8 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
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

func ReplaceInt16(slice []int16, old, new int16) []int16 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapInt16(slice []int16, f func(int16) int16) []int16 {
	for _, item := range slice {
		ReplaceInt16(slice, item, f(item))
	}
	return slice
}

func ReplaceAllInt16(slice []int16, old, new int16) []int16 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
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

func ReplaceInt32(slice []int32, old, new int32) []int32 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapInt32(slice []int32, f func(int32) int32) []int32 {
	for _, item := range slice {
		ReplaceInt32(slice, item, f(item))
	}
	return slice
}

func ReplaceAllInt32(slice []int32, old, new int32) []int32 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
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

func ReplaceInt64(slice []int64, old, new int64) []int64 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapInt64(slice []int64, f func(int64) int64) []int64 {
	for _, item := range slice {
		ReplaceInt64(slice, item, f(item))
	}
	return slice
}

func ReplaceAllInt64(slice []int64, old, new int64) []int64 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
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

func ReplaceFloat32(slice []float32, old, new float32) []float32 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapFloat32(slice []float32, f func(float32) float32) []float32 {
	for _, item := range slice {
		ReplaceFloat32(slice, item, f(item))
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

func ReplaceFloat64(slice []float64, old, new float64) []float64 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapFloat64(slice []float64, f func(float64) float64) []float64 {
	for _, item := range slice {
		ReplaceFloat64(slice, item, f(item))
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

func ReplaceString(slice []string, old, new string) []string {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapString(slice []string, f func(string) string) []string {
	for _, item := range slice {
		ReplaceString(slice, item, f(item))
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

func ReplaceUint(slice []uint, old, new uint) []uint {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapUint(slice []uint, f func(uint) uint) []uint {
	for _, item := range slice {
		ReplaceUint(slice, item, f(item))
	}
	return slice
}

func ReplaceAllUint(slice []uint, old, new uint) []uint {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
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

func ReplaceUint8(slice []uint8, old, new uint8) []uint8 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapUint8(slice []uint8, f func(uint8) uint8) []uint8 {
	for _, item := range slice {
		ReplaceUint8(slice, item, f(item))
	}
	return slice
}

func ReplaceAllUint8(slice []uint8, old, new uint8) []uint8 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
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

func ReplaceUint16(slice []uint16, old, new uint16) []uint16 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapUint16(slice []uint16, f func(uint16) uint16) []uint16 {
	for _, item := range slice {
		ReplaceUint16(slice, item, f(item))
	}
	return slice
}

func ReplaceAllUint16(slice []uint16, old, new uint16) []uint16 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
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

func ReplaceUint32(slice []uint32, old, new uint32) []uint32 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapUint32(slice []uint32, f func(uint32) uint32) []uint32 {
	for _, item := range slice {
		ReplaceUint32(slice, item, f(item))
	}
	return slice
}

func ReplaceAllUint32(slice []uint32, old, new uint32) []uint32 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
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

func ReplaceUint64(slice []uint64, old, new uint64) []uint64 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapUint64(slice []uint64, f func(uint64) uint64) []uint64 {
	for _, item := range slice {
		ReplaceUint64(slice, item, f(item))
	}
	return slice
}

func ReplaceAllUint64(slice []uint64, old, new uint64) []uint64 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
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

func ReplaceUintptr(slice []uintptr, old, new uintptr) []uintptr {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapUintptr(slice []uintptr, f func(uintptr) uintptr) []uintptr {
	for _, item := range slice {
		ReplaceUintptr(slice, item, f(item))
	}
	return slice
}

func ReplaceAllUintptr(slice []uintptr, old, new uintptr) []uintptr {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
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

func ReplaceComplex64(slice []complex64, old, new complex64) []complex64 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapComplex64(slice []complex64, f func(complex64) complex64) []complex64 {
	for _, item := range slice {
		ReplaceComplex64(slice, item, f(item))
	}
	return slice
}

func ReplaceAllComplex64(slice []complex64, old, new complex64) []complex64 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
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

func ReplaceComplex128(slice []complex128, old, new complex128) []complex128 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}

func MapComplex128(slice []complex128, f func(complex128) complex128) []complex128 {
	for _, item := range slice {
		ReplaceComplex128(slice, item, f(item))
	}
	return slice
}

func ReplaceAllComplex128(slice []complex128, old, new complex128) []complex128 {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
}
