package checksum

import (
	"crypto/rand"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha256(t *testing.T) {
	var body [100]byte
	if _, err := io.ReadFull(rand.Reader, body[:]); err != nil {
		assert.NoError(t, err)
	}

	body1, fail := Sha256Verify(body[:])
	assert.Equal(t, body[:], body1)
	assert.False(t, fail)

	sum := Sha256(body[:])
	assert.Len(t, sum, sha256Size)

	data := append(body[:], sum...)
	body2, ok := Sha256Verify(data)
	assert.True(t, ok)
	assert.Equal(t, body[:], body2)
}
