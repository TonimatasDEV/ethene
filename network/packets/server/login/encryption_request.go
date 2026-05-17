package login

import (
	"crypto/rand"
	"ethene/network/buffers"
	"fmt"
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

func GenerateVerifyToken() ([]byte, error) {
	nonce := make([]byte, 4)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random nonce: %w", err)
	}
	return nonce, nil
}
