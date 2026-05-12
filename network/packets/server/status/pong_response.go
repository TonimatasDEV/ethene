package status

import "ethene/network/util"

type PongResponse struct {
	Timestamp int64
}

func (p PongResponse) Id() int32 {
	return 1
}

func (p PongResponse) Marshal(buffer util.NetworkBuffer) {
	buffer.WriteLong(p.Timestamp)
}
