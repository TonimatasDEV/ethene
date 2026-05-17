package server

type MinecraftServer struct {
	PlayerList *PlayerList
}

func NewMinecraftServer() *MinecraftServer {
	server := new(MinecraftServer)
	server.PlayerList = NewPlayerList()
	return server
}
