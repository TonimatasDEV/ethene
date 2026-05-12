package status

import (
	"ethene/network/buffers"
)

type PingRequest struct {
	Timestamp int64
}

func (p *PingRequest) Unmarshal(buffer buffers.NetworkBuffer) error {
	p.Timestamp = buffer.ReadLong()
	return nil
}
