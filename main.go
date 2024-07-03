package main

import (
	"fmt"
	"github.com/chenyouwei3/Distributed-File/p2p"
	"log"
)

func OnPeer(peer p2p.Peer) error {
	fmt.Println("to doing something on transport")
	return nil
}

// 1.30.42
func main() {
	tcpOpts := p2p.TCPTransportOps{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
		//OnPeer: func(peer p2p.Peer) error {
		//	return fmt.Errorf("failed the onpeer func")
		//},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v", msg)
		}
	}()
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
