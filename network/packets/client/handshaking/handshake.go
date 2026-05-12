package handshaking

import (
	"ethene/network/buffers"
)

type HandshakePacket struct {
	Version    int32
	ServerName string
	Port       uint16 // TODO: Check that and Short
	State      int32
}

func (p *HandshakePacket) Unmarshal(buffer buffers.NetworkBuffer) error {
	p.Version, _ = buffer.ReadVarInt()
	p.ServerName = buffer.ReadString()
	var port = uint16(buffer.ReadShort())
	p.Port = port
	value, err := buffer.ReadVarInt()
	p.State = value
	return err
}
