package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	opt := TCPTransportOps{
		ListenAddr:    listenAddr,
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}
	tr := NewTCPTransport(opt)
	// 断言 ListenAddress 是否正确设置为 listenAddr
	assert.Equal(t, tr.ListenAddr, listenAddr)

	tr.ListenAndAccept()
	assert.Nil(t, tr.ListenAndAccept())
}
