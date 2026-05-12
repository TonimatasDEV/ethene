package packets

import "ethene/network/buffers"

type ClientPacket interface {
	Unmarshal(buffer buffers.NetworkBuffer) error
}

type ServerPacket interface {
	Id() int32
	Marshal(buffer buffers.NetworkBuffer)
}
