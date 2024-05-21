package main

import (
    "github.com/petspalace/orca"
)

func main() {
    t := orca.NewTerritory([16]byte{0})

    d := orca.NewDiscovery(t)
    d.Listen()
}
