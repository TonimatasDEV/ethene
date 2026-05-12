package status

import (
	"ethene/network/util"
)

type PingRequest struct {
	Timestamp int64
}

func (p *PingRequest) Unmarshal(buffer util.NetworkBuffer) error {
	p.Timestamp = buffer.ReadLong()
	return nil
}
