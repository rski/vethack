package vethack

import (
	"unsafe"
)

// IntoUnsafePointer takes a uintptr and converts it into an unsafe.Pointer,
// which is really bad and something you typically absolutely must not do.
func IntoUnsafePointer(v uintptr) unsafe.Pointer {
	return unsafe.Pointer(v)
}
