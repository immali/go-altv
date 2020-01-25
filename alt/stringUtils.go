package alt

import "unsafe"
import "C"

func StringToCCharPtr(str string) uintptr {
	cStr := C.CString(str)
	return (uintptr)(unsafe.Pointer(&cStr))
}
