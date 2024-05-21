package orca

type Peer struct {
	LastSeen int
}

func NewPeer() Peer {
	return Peer{}
}
