package alt

import (
	"fmt"
	"syscall"
)

type AltLib struct {
	lib  *syscall.DLL
	core uintptr

	coreLogInfo *syscall.Proc

	getSdkVersion *syscall.Proc

	scriptRuntimeCreate   *syscall.Proc
	scriptRuntimeRegister *syscall.Proc

	implCreate *syscall.Proc

	eventGetType                *syscall.Proc
	eventGetPlayerConnectTarget *syscall.Proc

	vector3Create *syscall.Proc

	playerGetRefFromEvent *syscall.Proc
	playerGetFromRef      *syscall.Proc
	playerGetID           *syscall.Proc
	playerSpawn           *syscall.Proc
	playerSetModel        *syscall.Proc
	playerGetName         *syscall.Proc

	resourceGetPath *syscall.Proc
	resourceGetMain *syscall.Proc
	resourceGetName *syscall.Proc

	stringViewGetData          *syscall.Proc
	stringViewAssignStringView *syscall.Proc
	stringViewCreate           *syscall.Proc

	stringAssignString *syscall.Proc
	stringCreate       *syscall.Proc
	stringGetData      *syscall.Proc
}

func (l *AltLib) LogInfo(format string, args ...interface{}) {
	l.coreLogInfo.Call(l.core, StringToCCharPtr(fmt.Sprintf(format, args...)))
}

func (l *AltLib) GetSDKVersion() int {
	res, _, _ := l.getSdkVersion.Call()
	return int(res)
}

func (l *AltLib) SetCore(core uintptr) {
	l.core = core
}

func NewAltLib(libPath string) *AltLib {
	lib := syscall.MustLoadDLL(libPath)

	altLib := &AltLib{
		lib: lib,

		coreLogInfo: lib.MustFindProc("alt_ICore_LogInfo"),

		getSdkVersion: lib.MustFindProc("alt_GetSDKVersion"),

		scriptRuntimeCreate:   lib.MustFindProc("alt_CAPIScriptRuntime_Create"),
		scriptRuntimeRegister: lib.MustFindProc("alt_ICore_RegisterScriptRuntime"),

		implCreate: lib.MustFindProc("alt_CAPIResource_Impl_Create"),

		eventGetType:                lib.MustFindProc("alt_CEvent_GetType"),
		eventGetPlayerConnectTarget: lib.MustFindProc("alt_CPlayerConnectEvent_GetTarget"),

		vector3Create: lib.MustFindProc("alt_Vector_float_3_PointLayout_Create_1"),

		playerGetID:           lib.MustFindProc("alt_IPlayer_GetID"),
		playerGetRefFromEvent: lib.MustFindProc("alt_CPlayerConnectEvent_GetTarget"),
		playerGetFromRef:      lib.MustFindProc("alt_RefBase_RefStore_IPlayer_Get"),
		playerSpawn:           lib.MustFindProc("alt_IPlayer_Spawn"),
		playerSetModel:        lib.MustFindProc("alt_IPlayer_SetModel"),
		playerGetName:         lib.MustFindProc("alt_IPlayer_GetName"),

		resourceGetPath: lib.MustFindProc("alt_IResource_GetPath"),
		resourceGetMain: lib.MustFindProc("alt_IResource_GetMain"),
		resourceGetName: lib.MustFindProc("alt_IResource_GetName"),

		stringViewGetData:          lib.MustFindProc("alt_StringView_GetData"),
		stringViewAssignStringView: lib.MustFindProc("alt_StringView_Assign_StringViewRefRef"),
		stringViewCreate:           lib.MustFindProc("alt_StringView_Create_2"),

		stringAssignString: lib.MustFindProc("alt_String_Assign_StringRefRef"),
		stringGetData:      lib.MustFindProc("alt_String_GetData"),
		stringCreate:       lib.MustFindProc("alt_String_Create_3"),
	}

	return altLib
}
