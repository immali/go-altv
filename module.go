package goaltv

import (
	"fmt"

	"github.com/immali/go-altv/alt"
)

type Module struct {
	lib       *alt.AltLib
	sr        *alt.ScriptRuntime
	rpc       *rpcServer
	resources map[string]*alt.Resource
}

func (m *Module) LogInfo(format string, args ...interface{}) {
	m.lib.LogInfo(fmt.Sprintf(format, args...))
}

func (m *Module) GetVersion() int {
	return m.lib.GetSDKVersion()
}

func (m *Module) Start() bool {
	m.sr = alt.NewScriptRuntime(m.lib, m.rpc.onEvent)
	res := m.sr.Register("go")

	if res {
		m.rpc = newRPCServer()
		m.rpc.listen()
	}

	return res
}

func (m *Module) SetCore(core uintptr) {
	m.lib.SetCore(core)
}

func NewModule() *Module {
	return &Module{
		lib: alt.NewAltLib(".\\altv-capi-server.dll"),
	}
}
