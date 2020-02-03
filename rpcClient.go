package goaltv

import (
	"fmt"
	"net"

	"github.com/cenkalti/rpc2"
)

type rpcClient struct {
	resID                   int
	rpc                     *rpc2.Client
	onPlayerConnectHandlers []func(p *Player)
	onResourceStartEvent    []func()
}

func (rc *rpcClient) run() {
	go rc.rpc.Run()
}

func (rc *rpcClient) serverHandleEvent(event string, handleArgs HandleArgs) bool {
	args := RPCHandleEventArgs{
		Event: event,
		ID:    rc.resID,
		Args:  handleArgs,
	}
	var reply RPCReply
	err := rc.rpc.Call(RPCHandleEvent, args, &reply)

	if err != nil {
		fmt.Println(err)
	}

	return bool(reply)
}

func (rc *rpcClient) registerEvent(event string, fn interface{}) bool {
	var callServer bool = false
	switch event {
	case RPCClientPlayerConnectedEvent:
		{
			callServer = len(rc.onPlayerConnectHandlers) == 0
			rc.onPlayerConnectHandlers = append(rc.onPlayerConnectHandlers, fn.(func(p *Player)))
			break
		}
	case RPCClientResourceStartEvent:
		{
			callServer = len(rc.onResourceStartEvent) == 0
			rc.onResourceStartEvent = append(rc.onResourceStartEvent, fn.(func()))
			break
		}
	}

	if !callServer {
		return false
	}

	var reply RPCReply
	args := RPCServerRegisterEventArgs{
		Event: event,
		ID:    rc.resID,
	}

	rc.rpc.Call(RPCServerRegisterEvent, args, &reply)
	return bool(reply)
}

func (rc *rpcClient) attachHandlers() {
	rpc := rc.rpc

	rpc.Handle(RPCHandleEvent, func(c *rpc2.Client, args *RPCHandleEventArgs, reply *RPCReply) error {
		switch args.Event {
		case RPCClientPlayerConnectedEvent:
			{
				p := NewPlayerFromConnectArgs(rc, args.Args)
				for _, fn := range rc.onPlayerConnectHandlers {
					fn(p)
				}
				break
			}
		case RPCClientResourceStartEvent:
			{
				for _, fn := range rc.onResourceStartEvent {
					fn()
				}
			}
		}
		*reply = RPCReply(true)
		return nil
	})
}

func newRPCClient(resID int) (*rpcClient, error) {
	conn, err := net.Dial("tcp", ":5000")

	if err != nil {
		return nil, err
	}

	rpc := rpc2.NewClient(conn)
	rc := &rpcClient{
		resID:                   resID,
		rpc:                     rpc,
		onPlayerConnectHandlers: make([]func(p *Player), 0),
		onResourceStartEvent:    make([]func(), 0),
	}

	rc.attachHandlers()

	return rc, nil
}
