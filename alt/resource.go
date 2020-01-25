package alt

type Resource struct {
	lib *AltLib
	ptr uintptr
}

func (r *Resource) GetMain() string {
	stringViewPtr, _, _ := r.lib.resourceGetMain.Call(r.ptr)
	stringView := NewStringViewFromPtr(r.lib, stringViewPtr)

	return stringView.GetData()
}

func (r *Resource) GetPath() string {
	stringViewPtr, _, _ := r.lib.resourceGetPath.Call(r.ptr)
	stringView := NewStringViewFromPtr(r.lib, stringViewPtr)

	return stringView.GetData()
}

func (r *Resource) GetName() string {
	stringViewPtr, _, _ := r.lib.resourceGetName.Call(r.ptr)
	stringView := NewStringViewFromPtr(r.lib, stringViewPtr)

	return stringView.GetData()
}

func NewResourceFromPtr(lib *AltLib, ptr uintptr) *Resource {
	return &Resource{
		lib: lib,
		ptr: ptr,
	}
}
