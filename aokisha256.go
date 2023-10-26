package aokisha256

import "github.com/PenisProtocol/aokisha256/internal"

// Hash returns the SHA256 hash of the given byte slice.
func Hash(target []byte) [32]byte {
	return internal.Hash(target)
}
