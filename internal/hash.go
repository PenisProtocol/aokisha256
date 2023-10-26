package internal

import "encoding/binary"

func Hash(target []byte) [32]byte {
	blocks := padIntoBlocks(target)

	prevState := constH

	for _, block := range blocks {
		prevState = processBlock(block, prevState)
	}

	var digest [32]byte
	binary.BigEndian.PutUint32(digest[0:4], prevState.a)
	binary.BigEndian.PutUint32(digest[4:8], prevState.b)
	binary.BigEndian.PutUint32(digest[8:12], prevState.c)
	binary.BigEndian.PutUint32(digest[12:16], prevState.d)
	binary.BigEndian.PutUint32(digest[16:20], prevState.e)
	binary.BigEndian.PutUint32(digest[20:24], prevState.f)
	binary.BigEndian.PutUint32(digest[24:28], prevState.g)
	binary.BigEndian.PutUint32(digest[28:32], prevState.h)

	return digest
}
