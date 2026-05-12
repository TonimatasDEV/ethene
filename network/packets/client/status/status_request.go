package status

import (
	"ethene/network/buffers"
)

type RequestStatus struct {
}

func (p *RequestStatus) Unmarshal(_ buffers.NetworkBuffer) error {
	return nil
}
