package login

import (
	"ethene/network/buffers"
)

type EncryptionResponse struct {
	SharedSecret []byte
	VerifyToken  []byte
}

func (p *EncryptionResponse) Unmarshal(buffer buffers.NetworkBuffer) error {
	p.SharedSecret = buffer.ReadBytes()
	p.VerifyToken = buffer.ReadBytes()
	return nil
}
