package util

import (
	"reflect"
	"unsafe"
)

func Int2Bool(v int) bool {
	return v != 0
}

func Bool2Int(v bool) int {
	if v {
		return 1
	}
	return 0
}

func SetSliceDataLen(pSlice, pData unsafe.Pointer, len int) {
	header := (*reflect.SliceHeader)(pSlice)
	header.Cap = len
	header.Len = len
	header.Data = uintptr(pData)
}

func GetZeroTermArrayLen(p unsafe.Pointer, size uintptr) int {
	var count uintptr
	//fmt.Println("size is", size)
	for {
		value := (*[999]byte)(unsafe.Pointer(uintptr(p) + (count * size)))
		//fmt.Println("value is", value)
		if allZero(value, size) {
			break
		}
		count++
	}

	return int(count)
}

func allZero(value *[999]byte, length uintptr) bool {
	for i := uintptr(0); i < length; i++ {
		if value[i] != byte(0) {
			return false
		}
	}
	return true
}
