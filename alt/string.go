package alt

import (
	"unsafe"
)

import "C"

type String struct {
	lib *AltLib
	ptr uintptr
}

func (s *String) GetData() string {
	charPtr, _, _ := s.lib.stringGetData.Call(s.ptr)
	cStr := (*C.char)(unsafe.Pointer(charPtr))

	return C.GoString(cStr)
}

func (s *String) AssignString(str string) {
	newString := NewStringFromString(s.lib, str)
	s.lib.stringAssignString.Call(s.ptr, newString.ptr)
}

func NewStringFromString(lib *AltLib, str string) *String {
	cStrPtr := (uintptr)(unsafe.Pointer(C.CString(str)))
	altStringPtr, _, _ := lib.stringCreate.Call(cStrPtr)

	return &String{
		lib: lib,
		ptr: altStringPtr,
	}
}

func NewStringFromPtr(lib *AltLib, ptr uintptr) *String {
	return &String{
		lib: lib,
		ptr: ptr,
	}
}
