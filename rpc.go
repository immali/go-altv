package goaltv

type RPCReply bool

type RPCServerRegisterEventArgs struct {
	Event string
	ID    int
}

type HandleArgs map[string]interface{}

type RPCHandleEventArgs struct {
	Event string
	ID    int
	Args  HandleArgs
}

var (
	RPCServerRegisterEvent        = "registerEvent"
	RPCServerSpawnPlayer          = "spawnPlayer"
	RPCServerPlayerSetModel       = "playerSetModel"
	RPCHandleEvent                = "handleEvent"
	RPCClientPlayerConnectedEvent = "playerConnected"
	RPCClientResourceStartEvent   = "resourceStarted"
)

func AltEventToRPCEvent(id int) string {
	switch id {
	case 1:
		return RPCClientPlayerConnectedEvent
	case 3:
		return RPCClientResourceStartEvent
	default:
		return ""
	}
}

// type RPCOnPlayerConnectEventArgs struct {
// 	ID   int
// 	Name string
// }

// type RPCSpawnPlayerEventArgs struct {
// 	ID  int
// 	Pos Vector3
// }

// type RPCPlayerSetModelEventArgs struct {
// 	ID    int
// 	Model uint
// }

// var (
// 	RPCNoneEvent            = "none"
// 	RPCRegisterEvent        = "registerEvent"
// 	RPCOnPlayerConnectEvent = "onPlayerConnect"
// 	RPCOnResourceStartEvent = "onResourceStart"
// 	RPCOnResourceStopEvent  = "onResourceStop"
// 	RPCSpawnPlayerEvent     = "rpcSpawnPlayer"
// 	RPCPlayerSetModelEvent  = "rpcPlayerSetModel"
// )

// // AltEventIDToRPCEvent - translates alt eventID to string
// func AltEventIDToRPCEvent(altEventID int) string {
// 	switch altEventID {
// 	case 1:
// 		return RPCOnPlayerConnectEvent
// 	case 3:
// 		return RPCOnResourceStartEvent
// 	case 4:
// 		return RPCOnResourceStopEvent
// 	default:
// 		return RPCNoneEvent
// 	}
// }
