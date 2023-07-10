package checksum

import (
	"bytes"
	"crypto/sha256"
)

// Sha256 returns the first 4 bytes of the sha256 checksum of data.
func Sha256(data []byte) []byte {
	b := sha256.Sum256(data)
	b = sha256.Sum256(b[:])
	return b[:4]
}

// Sha256Verify returns the body and whether the checksum of data is correct.
func Sha256Verify(data []byte) ([]byte, bool) {
	if len(data) < 4 {
		return nil, false
	}

	body, sum := data[:len(data)-4], data[len(data)-4:]
	return body, bytes.Equal(Sha256(body), sum)
}
