package keyexpansion

import "github.com/jackloughran/aes-go/shared"

func rotword(word uint32) uint32 {
	// a0, a1, a2, a3 -> a1, a2, a3, a0
	return word>>24 | word<<8
}

func subword(word uint32) uint32 {
	a0 := uint32(shared.SboxForward[(word&0xff000000)>>24]) << 24
	a1 := uint32(shared.SboxForward[(word&0x00ff0000)>>16]) << 16
	a2 := uint32(shared.SboxForward[(word&0x0000ff00)>>8]) << 8
	a3 := uint32(shared.SboxForward[(word & 0x000000ff)])

	return a0 | a1 | a2 | a3
}

func KeyExpansion(key []uint32, nK, nR int) [][]uint32 {
	w := make([]uint32, 4*(nR+1))

	for i := 0; i < nK; i++ {
		w[i] = key[i]
	}

	for i := nK; i < 4*(nR+1); i++ {
		temp := w[i-1]
		if i%nK == 0 {
			temp = rotword(subword(temp)) ^ shared.Rcon[i/nK-1]
		}

		w[i] = w[i-nK] ^ temp
	}

	encKeys := make([][]uint32, nR+1)
	for i := 0; i < nR+1; i++ {
		encKeys[i] = make([]uint32, 4)
		encKeys[i][0] = w[4*i]
		encKeys[i][1] = w[1+4*i]
		encKeys[i][2] = w[2+4*i]
		encKeys[i][3] = w[3+4*i]
	}

	return encKeys
}
