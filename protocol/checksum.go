package protocol

import (
	"bytes"
	"crypto/sha256"
)

func Sha256(data []byte) []byte {
	b := sha256.Sum256(data)
	b = sha256.Sum256(b[:])
	return b[:4]
}

func Sha256Verify(data []byte) ([]byte, bool) {
	if len(data) < 4 {
		return nil, false
	}

	body, sum := data[:len(data)-4], data[len(data)-4:]
	return body, bytes.Equal(Sha256(body), sum)
}
