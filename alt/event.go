package alt

type Event struct {
	lib *AltLib
	ptr uintptr
}

func (e *Event) GetType() int {
	typePtr, _, _ := e.lib.eventGetType.Call(e.ptr)
	return int(typePtr)
}

func (e *Event) GetPlayerConnectTarget() *Player {
	return NewPlayerByEvent(e.lib, e.ptr)
}

func NewEvent(lib *AltLib, altEvent uintptr) *Event {
	return &Event{
		lib: lib,
		ptr: altEvent,
	}
}
