package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	opt := TCPTransportOps{
		ListenAddr:    ":3000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}
	tr := NewTCPTransport(opt)
	// 断言 ListenAddress 是否正确设置为 listenAddr
	assert.Equal(t, tr.ListenAddr, ":3000")
	// 调用 ListenAndAccept 方法并断言其返回值为 nil
	assert.Nil(t, tr.ListenAndAccept())
}
