package subperm

import (
	"encoding/binary"

	"github.com/jackloughran/aes-go/shared"
)

func SubBytes(state []uint32) []uint32 {
	result := make([]uint32, len(state))
	for i, word := range state {
		b0 := uint32(shared.SboxForward[word>>24]) << 24
		b1 := uint32(shared.SboxForward[word>>16&0x000000ff]) << 16
		b2 := uint32(shared.SboxForward[word>>8&0x000000ff]) << 8
		b3 := uint32(shared.SboxForward[word&0x000000ff])

		result[i] = b0 | b1 | b2 | b3
	}

	return result
}

func ShiftRows(state []uint32) []uint32 {
	result := make([]uint32, 4)
	result[0] = state[0]
	result[1] = state[1]<<8 | state[1]>>24
	result[2] = state[2]<<16 | state[2]>>16
	result[3] = state[3]<<24 | state[3]>>8

	return result
}

func MixColumns(state []uint32) []uint32 {
	cols := shared.FlipRowsAndColumns(state)
	result := make([]uint32, 4)
	for i, col := range cols {
		words := make([]byte, 4)
		binary.BigEndian.PutUint32(words, col)
		d0 := shared.FiniteFieldMultiply(0x02, words[0]) ^ shared.FiniteFieldMultiply(0x3, words[1]) ^ words[2] ^ words[3]
		d1 := words[0] ^ shared.FiniteFieldMultiply(0x02, words[1]) ^ shared.FiniteFieldMultiply(0x03, words[2]) ^ words[3]
		d2 := words[0] ^ words[1] ^ shared.FiniteFieldMultiply(0x02, words[2]) ^ shared.FiniteFieldMultiply(0x03, words[3])
		d3 := shared.FiniteFieldMultiply(0x03, words[0]) ^ words[1] ^ words[2] ^ shared.FiniteFieldMultiply(0x02, words[3])

		result[i] = binary.BigEndian.Uint32([]byte{d0, d1, d2, d3})
	}

	return shared.FlipRowsAndColumns(result)
}

func AddRoundKey(state, key []uint32) []uint32 {
	result := make([]uint32, 4)
	for i, w := range state {
		result[i] = w ^ key[i]
	}

	return result
}
