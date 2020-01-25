package alt

// #include "c/creationInfo.c"
import (
	"C"
)
import "unsafe"

type CreationInfo struct {
	lib      *AltLib
	ptr      uintptr
	infoType *String
}

func (c *CreationInfo) getType() *String {
	// ptr := unsafe.Pointer(c.ptr)
	// infoPtr := (*C.alt_IResource_CreationInfo)(unsafe.Pointer(&c.ptr))
	cAltString := C.alt_IResource_CreationInfo_GetType(C.uintptr_t(c.ptr))

	altStringPtr := (uintptr)(unsafe.Pointer(cAltString))
	return NewStringFromPtr(c.lib, altStringPtr)
}

func (c *CreationInfo) GetType() string {
	return c.infoType.GetData()
}

func (c *CreationInfo) SetType(str string) {
	c.infoType.AssignString(str)
}

func NewCreationInfoFromPtr(lib *AltLib, ptr uintptr) *CreationInfo {
	creationInfo := &CreationInfo{
		lib: lib,
		ptr: ptr,
	}

	creationInfo.infoType = creationInfo.getType()

	return creationInfo
}
