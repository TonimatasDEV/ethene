package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"errors"
	"fmt"

	"golang.org/x/text/encoding/charmap"
)

var (
	ErrKeyGenerationFailed = errors.New("key pair generation failed")
	ErrDigestFailed        = errors.New("digest data failed")
	ErrDecryptionFailed    = errors.New("decryption failed")
	ErrCipherCreation      = errors.New("cipher creation failed")
)

type KeyPair struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func GenerateKeyPair() (*KeyPair, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrKeyGenerationFailed, err)
	}
	return &KeyPair{privateKey, &privateKey.PublicKey}, nil
}

func DigestData(data string, publicKey *rsa.PublicKey, secretKey []byte) ([]byte, error) {
	encoder := charmap.ISO8859_1.NewEncoder()
	dataBytes, err := encoder.Bytes([]byte(data))
	if err != nil {
		return nil, fmt.Errorf("failed to encode string to ISO_8859_1: %w", err)
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %w", err)
	}

	totalLen := len(dataBytes) + len(publicKeyBytes) + len(secretKey)
	buffer := make([]byte, 0, totalLen)
	buffer = append(buffer, dataBytes...)
	buffer = append(buffer, publicKeyBytes...)
	buffer = append(buffer, secretKey...)

	h := sha1.New()
	if _, err := h.Write(buffer); err != nil {
		return nil, fmt.Errorf("failed to write to hash: %w", err)
	}

	return h.Sum(nil), nil
}

func DecryptByteToSecretKey(privateKey *rsa.PrivateKey, encryptedBytes []byte) ([]byte, error) {
	decrypted, err := DecryptUsingKey(privateKey, encryptedBytes)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

func DecryptUsingKey(privateKey *rsa.PrivateKey, encryptedBytes []byte) ([]byte, error) {
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedBytes)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDecryptionFailed, err)
	}
	return decrypted, nil
}

func SetupCipher(mode int, key []byte) (cipher.Block, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCipherCreation, err)
	}
	return block, nil
}

func GetCipher(key []byte) (cipher.Stream, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, fmt.Errorf("invalid key size for AES: %d bytes (must be 16, 24, or 32)", len(key))
	}

	iv := key

	stream := cipher.NewCFBDecrypter(block, iv)
	return stream, nil
}

func DecryptCFB8(key []byte, ciphertext []byte) ([]byte, error) {
	stream, err := GetCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}

func SecretKeyToBytes(secretKey interface{}) ([]byte, error) {
	if bytes, ok := secretKey.([]byte); ok {
		return bytes, nil
	}
	return nil, errors.New("unsupported secret key type")
}

func PublicKeyToBytes(pubKey *rsa.PublicKey) ([]byte, error) {
	if pubKey == nil {
		return nil, fmt.Errorf("public key is nil")
	}

	bytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key to PKIX format: %w", err)
	}

	return bytes, nil
}
