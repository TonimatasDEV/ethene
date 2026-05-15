package login

import "ethene/network/buffers"

type Disconnect struct {
}

func (p Disconnect) Id() int32 {
	return 0
}

func (p Disconnect) Marshal(_ buffers.NetworkBuffer) {

}
