package alt

import (
	"unsafe"
)

import "C"

type StringView struct {
	lib *AltLib
	ptr uintptr
}

func (s *StringView) GetData() string {
	charPtr, _, _ := s.lib.stringViewGetData.Call(s.ptr)
	cString := (*C.char)(unsafe.Pointer(charPtr))

	return C.GoString(cString)
}

func (s *StringView) AssignString(str string) {
	cStr := (uintptr)(unsafe.Pointer(C.CString(str)))
	newStringView, _, _ := s.lib.stringViewCreate.Call(cStr)

	s.lib.stringViewAssignStringView.Call(s.ptr, newStringView)
}

func NewStringViewFromPtr(lib *AltLib, ptr uintptr) *StringView {
	return &StringView{
		lib: lib,
		ptr: ptr,
	}
}

func NewStringViewFromString(lib *AltLib, str string) *StringView {
	cStr := (uintptr)(unsafe.Pointer(C.CString(str)))
	newStringView, _, _ := lib.stringViewCreate.Call(cStr)

	return &StringView{
		lib: lib,
		ptr: newStringView,
	}
}
