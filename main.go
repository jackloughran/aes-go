package main

import (
	"fmt"

	"github.com/jackloughran/aes-go/keyexpansion"
	"github.com/jackloughran/aes-go/subperm"
)

func EncryptBlock128(input, key []uint32) []uint32 {
	keys := keyexpansion.KeyExpansion(key, 4, 10)

	state := input

	state = subperm.AddRoundKey(state, keys[0])

	for round := 1; round < 9; round++ {
		state = subperm.SubBytes(state)
		state = subperm.ShiftRows(state)
		state = subperm.MixColumns(state)
		state = subperm.AddRoundKey(state, keys[round])
	}

	state = subperm.SubBytes(state)
	state = subperm.ShiftRows(state)
	state = subperm.MixColumns(state)

	return state
}

func main() {
	fmt.Printf("%x\n", EncryptBlock128([]uint32{
		0x00000000,
		0x11111111,
		0x22222222,
		0x33333333,
	}, []uint32{
		0x01010101,
		0x02020202,
		0x03030303,
		0x04040404,
	}))
}
