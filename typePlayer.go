package goaltv

// Player - Resource.Player Object
type Player struct {
	rpc  *rpcClient
	ID   int
	Name string
	// Pos Vector3
}

// Spawn - spawns player at given location
func (p *Player) Spawn(pos Vector3) bool {
	args := HandleArgs{
		"ID": p.ID,
		"X":  pos.X,
		"Y":  pos.Y,
		"Z":  pos.Z,
	}

	return p.rpc.serverHandleEvent(RPCServerSpawnPlayer, args)
}

func (p *Player) SetModel(model uint) {
	args := HandleArgs{
		"ID":    p.ID,
		"Model": model,
	}

	p.rpc.serverHandleEvent(RPCServerPlayerSetModel, args)
}

func NewPlayerFromConnectArgs(rpc *rpcClient, args HandleArgs) *Player {
	return &Player{
		rpc:  rpc,
		ID:   args["ID"].(int),
		Name: args["Name"].(string),
	}
}
