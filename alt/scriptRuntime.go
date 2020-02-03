package alt

import (
	"syscall"
)

type ScriptRuntime struct {
	lib     *AltLib
	ptr     uintptr
	impl    []*ResourceImpl
	onEvent func(id int, event *Event)
}

func (sr *ScriptRuntime) createImpl(altScriptRuntime, altResource uintptr) uintptr {

	id := len(sr.impl)
	impl := NewResourceImpl(sr.lib, altResource, id, sr.onNotifyEvent)
	sr.impl = append(sr.impl, impl)

	return impl.ptr
}

func (sr *ScriptRuntime) destroyImpl() uintptr {
	return 0
}

func (sr *ScriptRuntime) onTick() uintptr {
	return 0
}

func (sr *ScriptRuntime) Register(moduleType string) bool {
	typeStringView := NewStringViewFromString(sr.lib, moduleType)
	res, _, _ := sr.lib.scriptRuntimeRegister.Call(sr.lib.core, typeStringView.ptr, sr.ptr)

	if int(res) == 1 {
		return true
	}

	return false
}

func (sr *ScriptRuntime) onNotifyEvent(resID int, event *Event) {
	sr.onEvent(resID, event)
}

func NewScriptRuntime(lib *AltLib, onEvent func(id int, event *Event)) *ScriptRuntime {
	sr := &ScriptRuntime{lib: lib, onEvent: onEvent}

	ptr, _, _ := lib.scriptRuntimeCreate.Call(
		syscall.NewCallback(sr.createImpl),
		syscall.NewCallback(sr.destroyImpl),
		syscall.NewCallback(sr.onTick))

	sr.ptr = ptr

	return sr
}
