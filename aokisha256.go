package aokisha256

import "github.com/PenisProtocol/aokisha256/internal"

func Hash(target []byte) [32]byte {
	return internal.Hash(target)
}
