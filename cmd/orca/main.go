package main

import (
    "github.com/petspalace/orca"
)

func main() {
    t := orca.NewDiscovery()
    t.Listen()
}
