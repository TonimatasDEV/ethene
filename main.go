package main

import (
	"ethene/network"
	"ethene/server"
)

func main() {
	server.NewMinecraftServer()
	network.InitReceiver()
}
