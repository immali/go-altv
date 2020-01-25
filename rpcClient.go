package goaltv

import (
	"net"

	"github.com/cenkalti/rpc2"
)

type rpcClient struct {
	resID                   int
	rpc                     *rpc2.Client
	onPlayerConnectHandlers []func(p *Player)
}

func (c *rpcClient) run() {
	go c.rpc.Run()
}

func (c *rpcClient) addOnPlayerConnectHandler(handler func(p *Player)) {
	callServer := len(c.onPlayerConnectHandlers) == 0
	c.onPlayerConnectHandlers = append(c.onPlayerConnectHandlers, handler)

	if callServer {
		var reply RPCReply
		args := RPCRegisterEventArgs{
			Event: RPCOnPlayerConnectEvent,
			ID:    c.resID,
		}
		c.rpc.Call(RPCRegisterEvent, &args, &reply)
	}
}

func newRPCClient(resID int) (*rpcClient, error) {
	conn, err := net.Dial("tcp", ":5000")

	if err != nil {
		return nil, err
	}

	c := &rpcClient{
		resID:                   resID,
		rpc:                     rpc2.NewClient(conn),
		onPlayerConnectHandlers: make([]func(p *Player), 0),
	}

	c.rpc.Handle(RPCOnPlayerConnectEvent, func(client *rpc2.Client, args *RPCOnPlayerConnectEventArgs, reply *RPCReply) error {
		for _, h := range c.onPlayerConnectHandlers {
			h(NewPlayerFromConnectArgs(c.rpc, args))
		}

		*reply = RPCReply(true)
		return nil
	})

	return c, nil
}
