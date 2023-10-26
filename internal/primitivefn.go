package internal

func rotr(x uint32, n uint) uint32 {
	return (x >> n) | (x << (32 - n))
}

func sigma0(x uint32) uint32 {
	return rotr(x, 2) ^ rotr(x, 13) ^ rotr(x, 22)
}

func sigma1(x uint32) uint32 {
	return rotr(x, 6) ^ rotr(x, 11) ^ rotr(x, 25)
}

func lsigma0(x uint32) uint32 {
	return rotr(x, 7) ^ rotr(x, 18) ^ (x >> 3)
}

func lsigma1(x uint32) uint32 {
	return rotr(x, 17) ^ rotr(x, 19) ^ (x >> 10)
}

func ch(x uint32, y uint32, z uint32) uint32 {
	return (x & y) ^ (^x & z)
}

func maj(x uint32, y uint32, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}
