package main

import (
	"crypto/aes"
	"encoding/binary"
	"testing"
)

func TestEncryptBlock128(t *testing.T) {
	plaintext := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
	}

	key := []byte{
		0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
	}

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		t.Error(err)
	}

	expected := make([]byte, len(plaintext))
	cipherBlock.Encrypt(expected, plaintext)

	actual := EncryptBlock128(byteToUint32(plaintext), byteToUint32(key))

	if !equal(byteToUint32(expected), actual) {
		t.Errorf("expected: %x, actual: %x", byteToUint32(expected), actual)
	}
}

func byteToUint32(in []byte) []uint32 {
	result := make([]uint32, len(in)/4)
	for i := 0; i < len(in)/4; i++ {
		result[i] = binary.BigEndian.Uint32(in[i*4 : (i+1)*4])
	}

	return result
}

func equal(a, b []uint32) bool {
	for i, x := range a {
		if b[i] != x {
			return false
		}
	}

	return true
}
