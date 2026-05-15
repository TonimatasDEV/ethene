package server

import (
	"ethene/network"
)

type MinecraftServer struct {
	PlayerList *PlayerList
}

func NewMinecraftServer() *MinecraftServer {
	server := new(MinecraftServer)
	server.PlayerList = NewPlayerList()

	network.InitReceiver()

	return server
}
