package login

import "ethene/network/buffers"

type SetCompression struct {
	Threshold int32
}

func (p *SetCompression) Id() int32 {
	return 3
}

func (p *SetCompression) Marshal(buffer buffers.NetworkBuffer) {
	buffer.WriteVarInt(p.Threshold)
}
