package goaltv

import (
	"fmt"
	"net"

	"github.com/cenkalti/rpc2"
	"github.com/immali/go-altv/alt"
)

type rpcServer struct {
	r *rpc2.Server
}

var listeners = make(map[int]map[string]*rpc2.Client)
var players = make(map[int]*alt.Player)

func (rs *rpcServer) onEvent(id int, altEvent *alt.Event) {
	var args RPCHandleEventArgs
	var reply RPCReply
	event := AltEventToRPCEvent(altEvent.GetType())

	fmt.Println("AltEvent: ", event)

	if _, exists := listeners[id]; !exists {
		return
	}

	if _, exists := listeners[id][event]; !exists {
		return
	}

	fmt.Println(fmt.Sprintf("Handle AltEvent %s for Resource %d", event, id))

	switch event {
	case RPCClientPlayerConnectedEvent:
		{
			p := altEvent.GetPlayerConnectTarget()
			playerID := p.GetID()
			args = RPCHandleEventArgs{
				Event: RPCClientPlayerConnectedEvent,
				ID:    id,
				Args: HandleArgs{
					"ID":   playerID,
					"Name": p.GetName(),
				},
			}

			if _, exists := players[playerID]; !exists {
				players[playerID] = p
			}
			break
		}
	case RPCClientResourceStartEvent:
		{
			args = RPCHandleEventArgs{
				Event: RPCClientResourceStartEvent,
				ID:    id,
				Args:  HandleArgs{},
			}
			break
		}
	default:
		{
			return
		}
	}

	fmt.Println("Calling resource")
	client, _ := listeners[id][event]
	err := client.Call(RPCHandleEvent, args, &reply)

	if err != nil {
		fmt.Println(err)
	}
}

func (rs *rpcServer) attachHandlers() {
	rpc := rs.r

	rpc.Handle(RPCServerRegisterEvent, func(c *rpc2.Client, args *RPCServerRegisterEventArgs, reply *RPCReply) error {
		if _, exists := listeners[args.ID]; !exists {
			listeners[args.ID] = make(map[string]*rpc2.Client)
		}

		if _, exists := listeners[args.ID][args.Event]; !exists {
			listeners[args.ID][args.Event] = c
		}

		*reply = RPCReply(true)
		return nil
	})

	rpc.Handle(RPCHandleEvent, func(c *rpc2.Client, args *RPCHandleEventArgs, reply *RPCReply) error {
		switch args.Event {
		case RPCServerSpawnPlayer:
			{
				id := args.Args["ID"].(int)
				pos := Vector3{
					X: args.Args["X"].(float64),
					Y: args.Args["Y"].(float64),
					Z: args.Args["Z"].(float64),
				}

				if player, exists := players[id]; exists {
					player.Spawn(pos.X, pos.Y, pos.Z, 0)
				}
				break
			}
		case RPCServerPlayerSetModel:
			{
				id := args.Args["ID"].(int)
				model := args.Args["Model"].(uint)

				if player, exists := players[id]; exists {
					player.SetModel(model)
				}
				break
			}
		}

		*reply = RPCReply(true)
		return nil
	})
}

func (rs *rpcServer) listen() {
	lis, _ := net.Listen("tcp", "127.0.0.1:5000")
	go rs.r.Accept(lis)

}

func newRPCServer() *rpcServer {
	r := rpc2.NewServer()
	rs := &rpcServer{r}
	rs.attachHandlers()

	return rs
}
