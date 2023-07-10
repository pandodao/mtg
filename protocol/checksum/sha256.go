package checksum

import (
	"bytes"
	"crypto/sha256"
)

const (
	sha256Size = 4
)

// Sha256 returns the first 4 bytes of the sha256 checksum of data.
func Sha256(data []byte) []byte {
	b := sha256.Sum256(data)
	b = sha256.Sum256(b[:])
	return b[:sha256Size]
}

// Sha256Verify returns the body and whether the checksum of data is correct.
func Sha256Verify(data []byte) ([]byte, bool) {
	if len(data) < sha256Size {
		return nil, false
	}

	body, sum := data[:len(data)-sha256Size], data[len(data)-sha256Size:]
	return body, bytes.Equal(Sha256(body), sum)
}
