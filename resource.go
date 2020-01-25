package goaltv

import (
	"flag"
)

type Resource struct {
	id  int
	rpc *rpcClient
}

func (r *Resource) OnPlayerConnect(handler func(p *Player)) {
	r.rpc.addOnPlayerConnectHandler(handler)
}

func NewResource() *Resource {
	id := flag.Int("id", -1, "Resource ID")
	flag.Parse()

	if *id < 0 {
		panic("WRONG RESOURCE ID")
	}

	rpc, err := newRPCClient(*id)

	if err != nil {
		panic("COULD NOT CONNECT TO SERVER-RESOURCE")
	}

	r := &Resource{
		id:  *id,
		rpc: rpc,
	}

	r.rpc.run()

	return r
}
