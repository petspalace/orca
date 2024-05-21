package orca

type Territory struct {
	PSK   [16]byte
	Peers []Peer
}

func NewTerritory(PSK [16]byte) Territory {
	return Territory{PSK: PSK}
}
