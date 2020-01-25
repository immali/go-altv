package goaltv

type RPCReply bool

type RPCRegisterEventArgs struct {
	Event string
	ID    int
}

type RPCOnPlayerConnectEventArgs struct {
	ID   int
	Name string
}

type RPCSpawnPlayerEventArgs struct {
	ID  int
	Pos Vector3
}

type RPCPlayerSetModelEventArgs struct {
	ID    int
	Model uint
}

var (
	RPCNoneEvent            = "none"
	RPCRegisterEvent        = "registerEvent"
	RPCOnPlayerConnectEvent = "onPlayerConnect"
	RPCOnResourceStartEvent = "onResourceStart"
	RPCOnResourceStopEvent  = "onResourceStop"
	RPCSpawnPlayerEvent     = "rpcSpawnPlayer"
	RPCPlayerSetModelEvent  = "rpcPlayerSetModel"
)

// AltEventIDToRPCEvent - translates alt eventID to string
func AltEventIDToRPCEvent(altEventID int) string {
	switch altEventID {
	case 1:
		return RPCOnPlayerConnectEvent
	case 3:
		return RPCOnResourceStartEvent
	case 4:
		return RPCOnResourceStopEvent
	default:
		return RPCNoneEvent
	}
}
