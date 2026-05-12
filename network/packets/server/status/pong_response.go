package status

import "ethene/network/buffers"

type PongResponse struct {
	Timestamp int64
}

func (p PongResponse) Id() int32 {
	return 1
}

func (p PongResponse) Marshal(buffer buffers.NetworkBuffer) {
	buffer.WriteLong(p.Timestamp)
}
