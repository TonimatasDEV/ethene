package server

import "ethene/world/entities"

type PlayerList struct {
	Players []entities.Player
}

func NewPlayerList() *PlayerList {
	return &PlayerList{
		Players: make([]entities.Player, 0),
	}
}

func (playerList *PlayerList) Add(player entities.Player) {
	playerList.Players = append(playerList.Players, player)
}
