package login

import (
	"crypto/rand"
	"ethene/network/buffers"
)

type EncryptionRequest struct {
	ServerID           string
	PublicKey          []byte
	VerifyToken        []byte
	ShouldAuthenticate bool
}

func (p EncryptionRequest) Id() int32 {
	return 1
}

func (p EncryptionRequest) Marshal(buffer buffers.NetworkBuffer) {
	buffer.WriteString(p.ServerID)
	buffer.WriteBytes(p.PublicKey)
	buffer.WriteBytes(p.VerifyToken)
	buffer.WriteBool(p.ShouldAuthenticate)
}

func GenerateVerifyToken(length int) ([]byte, error) {
	token := make([]byte, length)
	_, err := rand.Read(token)
	return token, err
}
