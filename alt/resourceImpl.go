package alt

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"syscall"
	"unsafe"
)

type ResourceImpl struct {
	lib         *AltLib
	ptr         uintptr
	resource    *Resource
	id          int
	notifyEvent func(id int, event *Event)
}

func (ri *ResourceImpl) makeClient(altResource, altCreationInfo, altFiles uintptr) uintptr {
	creationInfo := NewCreationInfoFromPtr(ri.lib, altCreationInfo)
	creationInfo.SetType("js")

	res := true
	return (uintptr)(unsafe.Pointer(&res))
}

func (ri *ResourceImpl) start(altResource uintptr) uintptr {
	pathToMainFile := path.Join(ri.resource.GetPath(), ri.resource.GetMain())
	idFlag := fmt.Sprintf("-id=%d", ri.id)
	cmd := exec.Command("go", "run", pathToMainFile, idFlag)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go cmd.Start()

	res := true
	return (uintptr)(unsafe.Pointer(&res))
}

func (ri *ResourceImpl) stop() uintptr { return 0 }
func (ri *ResourceImpl) onEvent(_, altEvent uintptr) uintptr {
	ri.notifyEvent(ri.id, NewEvent(ri.lib, altEvent))

	res := true
	return (uintptr)(unsafe.Pointer(&res))
}
func (ri *ResourceImpl) onTick() uintptr             { return 0 }
func (ri *ResourceImpl) onCreateBaseObject() uintptr { return 0 }
func (ri *ResourceImpl) onRemoveBaseObject() uintptr { return 0 }

func (ri *ResourceImpl) createImpl(altResource uintptr) {
	implPtr, _, _ := ri.lib.implCreate.Call(
		altResource,
		syscall.NewCallback(ri.makeClient),
		syscall.NewCallback(ri.start),
		syscall.NewCallback(ri.stop),
		syscall.NewCallback(ri.onEvent),
		syscall.NewCallback(ri.onTick),
		syscall.NewCallback(ri.onCreateBaseObject),
		syscall.NewCallback(ri.onRemoveBaseObject))

	ri.ptr = implPtr
}

func NewResourceImpl(lib *AltLib, altResource uintptr, id int, notifyEvent func(id int, event *Event)) *ResourceImpl {
	ri := &ResourceImpl{
		lib:         lib,
		resource:    NewResourceFromPtr(lib, altResource),
		id:          id,
		notifyEvent: notifyEvent,
	}

	ri.createImpl(altResource)
	return ri
}
