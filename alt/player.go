package alt

import (
	"unsafe"
)

type Player struct {
	ptr uintptr
	ref uintptr
	lib *AltLib
}

func (p *Player) GetID() int {
	idPtr, _, _ := p.lib.playerGetID.Call(p.ptr)
	return int(idPtr)
}

func (p *Player) GetName() string {
	stringViewPtr, _, _ := p.lib.playerGetName.Call(p.ptr)
	stringView := NewStringViewFromPtr(p.lib, stringViewPtr)

	return stringView.GetData()
}

func (p *Player) Spawn(X, Y, Z float64, delay int) {
	vec3 := NewVector3(p.lib, X, Y, Z)
	delayPtr := (uintptr)(unsafe.Pointer(&delay))

	p.lib.playerSpawn.Call(p.ptr, vec3.ptr, delayPtr)
}

func (p *Player) SetModel(model uint) {
	p.lib.playerSetModel.Call(p.ptr, (uintptr)(model))
}

func NewPlayerByEvent(lib *AltLib, event uintptr) *Player {
	playerRef, _, _ := lib.playerGetRefFromEvent.Call(event)
	playerPtr, _, _ := lib.playerGetFromRef.Call(playerRef)

	return &Player{
		lib: lib,
		ptr: playerPtr,
		ref: playerRef,
	}
}
