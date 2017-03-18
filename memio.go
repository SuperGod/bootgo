package kernel

import "unsafe"

func WriteMemory(ptr uintptr, pos int, v uint8) {
	addr := (*uint8)(unsafe.Pointer(ptr + uintptr(pos)))
	*addr = v
	return
}

func ReadMemory(ptr uintptr, pos int) uint8 {
	addr := (*uint8)(unsafe.Pointer(ptr + uintptr(pos)))
	return *addr
}
