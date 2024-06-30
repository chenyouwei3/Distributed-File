package p2p

import "errors"

var ErrInvalidHandshake = errors.New("handshake 无效")

type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
