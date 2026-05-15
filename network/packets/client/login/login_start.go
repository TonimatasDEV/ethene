package login

import (
	"ethene/network/buffers"

	"github.com/google/uuid"
)

type StartLogin struct {
	Name       string
	PlayerUUID uuid.UUID
}

func (p *StartLogin) Unmarshal(buffer buffers.NetworkBuffer) error {
	p.Name = buffer.ReadString()
	p.PlayerUUID = buffer.ReadUUID()
	return nil
}
