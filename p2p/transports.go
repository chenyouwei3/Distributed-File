package p2p

type Peer interface {
}

// transport is anything that handles the communitcation
// between the nodes in the network
type Transport interface {
	ListenAndAccept() error
}
