package internal

import "encoding/binary"

func padIntoBlocks(rawdata []byte) (blocks [][64]byte) {
	rawLen := len(rawdata)

	// calculate the number of blocks
	var targetBlockLen int
	if (rawLen+9)%64 == 0 {
		targetBlockLen = (rawLen + 9) / 64
	} else {
		targetBlockLen = (rawLen+9)/64 + 1
	}

	// allocate blocks
	blocks = make([][64]byte, targetBlockLen)

	// copy rawdata into blocks
	for i := 0; i < (rawLen/64)+1; i++ {
		var btm int
		if (i+1)*64 > rawLen {
			btm = rawLen
		} else {
			btm = (i + 1) * 64
		}
		srcSlice := rawdata[i*64 : btm]
		copy(blocks[i][:], srcSlice)
	}

	// append 0x80
	lastContBlockNum := rawLen / 64
	blocks[lastContBlockNum][rawLen%64] = 0x80

	// append length

	binary.BigEndian.PutUint64(blocks[targetBlockLen-1][56:64], uint64(rawLen*8))

	return blocks
}
