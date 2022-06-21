package vethack

import (
	"unsafe"
)

// DoIntoUnsafePtr takes a function, typically a syscall,
// that returns a uintptr and converts it into an unsafe.Pointer.
func IntoUnsafePointer(v uintptr) unsafe.Pointer {
	return unsafe.Pointer(v)
}
