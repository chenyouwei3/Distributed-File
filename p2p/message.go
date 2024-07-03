package p2p

import "net"

//知道数据从哪里来

type RPC struct {
	From    net.Addr //知道数据从哪里来
	Payload []byte
}
