package login

import "ethene/network/buffers"

type Disconnect struct {
	Reason string
}

func (p *Disconnect) Id() int32 {
	return 0
}

func (p *Disconnect) Marshal(buffer buffers.NetworkBuffer) {
	buffer.WriteString(p.Reason)
}
