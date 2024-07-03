package p2p

type Peer interface {
	Close() error
}

// transport is anything that handles the communitcation
// between the nodes in the network
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
