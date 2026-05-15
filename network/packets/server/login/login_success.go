package login

import "ethene/network/buffers"

type SuccessLogin struct {
}

func (p SuccessLogin) Id() int32 {
	return 2
}

func (p SuccessLogin) Marshal(_ buffers.NetworkBuffer) {

}
