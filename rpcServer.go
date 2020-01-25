package goaltv

import (
	"fmt"
	"net"

	"github.com/cenkalti/rpc2"
	"github.com/immali/go-altv/alt"
)

var players map[int]*alt.Player = make(map[int]*alt.Player)
var onPlayerConnect map[int]*rpc2.Client = make(map[int]*rpc2.Client)

type rpcServer struct {
	rpc *rpc2.Server
}

func (s *rpcServer) emitPlayerConnect(resID int, p *alt.Player) bool {
	client, listening := onPlayerConnect[resID]

	if !listening {
		return false
	}

	var reply RPCReply
	args := RPCOnPlayerConnectEventArgs{
		ID:   p.GetID(),
		Name: p.GetName(),
	}

	client.Call(RPCOnPlayerConnectEvent, &args, &reply)

	return bool(reply)
}

func (s *rpcServer) onEvent(id int, event *alt.Event) {
	switch AltEventIDToRPCEvent(event.GetType()) {
	case RPCOnPlayerConnectEvent:
		{
			p := event.GetPlayerConnectTarget()
			if _, exists := players[p.GetID()]; !exists {
				players[p.GetID()] = p
			}

			s.emitPlayerConnect(id, event.GetPlayerConnectTarget())
			break
		}
	}
}

func (s *rpcServer) listen() {
	lis, _ := net.Listen("tcp", "127.0.0.1:5000")
	go s.rpc.Accept(lis)
}

func newRPCServer() *rpcServer {
	s := &rpcServer{
		rpc: rpc2.NewServer(),
	}

	s.rpc.Handle(RPCRegisterEvent, func(client *rpc2.Client, args *RPCRegisterEventArgs, reply *RPCReply) error {
		fmt.Println(fmt.Sprintf("Registering event %s for Resource %d", args.Event, args.ID))

		if args.Event == RPCNoneEvent {
			*reply = RPCReply(true)
			return nil
		}

		switch args.Event {
		case RPCOnPlayerConnectEvent:
			{
				onPlayerConnect[args.ID] = client
				break
			}
		}

		*reply = RPCReply(true)
		return nil
	})

	s.rpc.Handle(RPCSpawnPlayerEvent, func(client *rpc2.Client, args *RPCSpawnPlayerEventArgs, reply *RPCReply) error {
		if p, exists := players[args.ID]; exists {
			p.Spawn(args.Pos.X, args.Pos.Y, args.Pos.Z, 10)
		}

		*reply = RPCReply(true)
		return nil
	})

	s.rpc.Handle(RPCPlayerSetModelEvent, func(_ *rpc2.Client, args *RPCPlayerSetModelEventArgs, reply *RPCReply) error {
		if p, exists := players[args.ID]; exists {
			p.SetModel(args.Model)
		}

		*reply = RPCReply(true)
		return nil
	})

	return s
}
