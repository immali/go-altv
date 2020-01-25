package alt

import "C"
import (
	"fmt"
	"unsafe"
)

type Vector3 struct {
	lib *AltLib
	ptr uintptr
}

func NewVector3(lib *AltLib, X, Y, Z float64) *Vector3 {
	x := (uintptr)(unsafe.Pointer(&X))
	y := (uintptr)(unsafe.Pointer(&Y))
	z := (uintptr)(unsafe.Pointer(&Z))

	// cX := C.float(X)
	// cY := C.float(Y)
	// cZ := C.float(Z)

	// x := (uintptr)(unsafe.Pointer(&cX))
	// y := (uintptr)(unsafe.Pointer(&cY))
	// z := (uintptr)(unsafe.Pointer(&cZ))

	altVector3Ptr, _, _ := lib.vector3Create.Call(x, y, z)

	fmt.Println("vector: ", altVector3Ptr)

	return &Vector3{
		lib: lib,
		ptr: altVector3Ptr,
	}
}
