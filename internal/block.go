package internal

import "encoding/binary"

func schedArr(block [64]byte) [64]uint32 {
	var sched [64]uint32
	for i := 0; i < 16; i++ {
		sched[i] = binary.BigEndian.Uint32(block[i*4 : (i+1)*4])
	}
	for i := 16; i < 64; i++ {
		sched[i] = lsigma1(sched[i-2]) + sched[i-7] + lsigma0(sched[i-15]) + sched[i-16]
	}
	return sched
}

type state struct {
	a, b, c, d, e, f, g, h uint32
}

func processBlock(block [64]byte, prevState state) (newState state) {
	sched := schedArr(block)

	a, b, c, d, e, f, g, h := prevState.a, prevState.b, prevState.c, prevState.d, prevState.e, prevState.f, prevState.g, prevState.h

	for i := 0; i < 64; i++ {
		t1 := h + sigma1(e) + ch(e, f, g) + constK[i] + sched[i]
		t2 := sigma0(a) + maj(a, b, c)
		h, g, f, e, d, c, b, a = g, f, e, d+t1, c, b, a, t1+t2
	}

	newState.a = prevState.a + a
	newState.b = prevState.b + b
	newState.c = prevState.c + c
	newState.d = prevState.d + d
	newState.e = prevState.e + e
	newState.f = prevState.f + f
	newState.g = prevState.g + g
	newState.h = prevState.h + h

	return newState
}
