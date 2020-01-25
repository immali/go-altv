package goaltv

import (
	"github.com/cenkalti/rpc2"
)

// Player - Resource.Player Object
type Player struct {
	rpc  *rpc2.Client
	ID   int
	Name string
	// Pos Vector3
}

// Spawn - spawns player at given location
func (p *Player) Spawn(pos Vector3) bool {
	var reply RPCReply
	args := RPCSpawnPlayerEventArgs{
		ID:  p.ID,
		Pos: pos,
	}

	p.rpc.Call(RPCSpawnPlayerEvent, &args, &reply)

	return bool(reply)
}

func (p *Player) SetModel(model uint) {
	var reply RPCReply
	args := RPCPlayerSetModelEventArgs{ID: p.ID, Model: model}

	p.rpc.Call(RPCPlayerSetModelEvent, &args, &reply)
}

func NewPlayerFromConnectArgs(rpc *rpc2.Client, args *RPCOnPlayerConnectEventArgs) *Player {
	return &Player{
		rpc:  rpc,
		ID:   args.ID,
		Name: args.Name,
	}
}
