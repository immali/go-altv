package alt

// #include "c/vector3.c"
import "C"

type Vector3 struct {
	lib *AltLib
	ptr uintptr
}

func NewVector3(lib *AltLib, X, Y, Z float64) *Vector3 {
	cX := C.float(X)
	cY := C.float(Y)
	cZ := C.float(Z)

	altVector3Ptr, _, _ := lib.vector3Create.Call()
	C.alt_Vector_float_3_PointLayout_Set(C.uintptr_t(altVector3Ptr), cX, cY, cZ)

	return &Vector3{
		lib: lib,
		ptr: altVector3Ptr,
	}
}
