package aokisha256_test

import (
	"crypto/sha256"
	"math/rand"
	"testing"

	"github.com/PenisProtocol/aokisha256"
)

func TestHash(t *testing.T) {
	tests := []struct {
		name   string
		target []byte
	}{
		{
			name:   "empty input",
			target: []byte{},
		},
		{
			name:   "short input",
			target: []byte("hello"),
		},
		{
			name:   "long input",
			target: []byte("The quick brown fox jumps over the lazy dog"),
		},
		{
			name:   "watashi no mai nanba- (excess 64 bytes)",
			target: []byte("Magic Number is 9774-8380-6896. It is advised to be kept secret."),
		},
		{
			name:   "watashi no mai nanba- (excess 56 bytes)",
			target: []byte("Magic Number is 9774-8380-6896. You must Keep it secret."),
		},
		{
			name:   "watashi no mai nanba- (less than 56 bytes)",
			target: []byte("Magic Number is 9774-8380-6896. You must keep it secret"),
		},
		{
			name:   "long input",
			target: []byte("The quick brown fox jumps over the lazy dog. The quick brown fox jumps over the lazy dog. The quick brown fox jumps over the lazy dog."),
		},

		{
			name: "ultra long input",
			target: func() []byte {
				// generate random value
				var target []byte
				for i := 0; i < 10000000; i++ {
					target = append(target, byte(rand.Intn(256)))
				}
				return target
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := sha256.Sum256(tt.target)
			got := aokisha256.Hash(tt.target)
			if got != expected {
				t.Errorf("Hash(%v) = %v, want %v", tt.target, got, expected)
			}
		})
	}
}
