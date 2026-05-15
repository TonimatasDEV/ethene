package entities

import (
	"ethene/network"

	"github.com/google/uuid"
)

type Player struct {
	UUID       uuid.UUID
	Name       string
	Connection *network.Connection
}
