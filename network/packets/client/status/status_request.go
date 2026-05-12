package status

import (
	"ethene/network/util"
)

type RequestStatus struct {
}

func (p *RequestStatus) Unmarshal(_ util.NetworkBuffer) error {
	return nil
}
