package packets

import "ethene/network/util"

type ClientPacket interface {
	Unmarshal(buffer util.NetworkBuffer) error
}

type ServerPacket interface {
	Id() int32
	Marshal(buffer util.NetworkBuffer)
}
