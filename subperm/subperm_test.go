package subperm

import (
	"crypto/aes"
	"testing"
)

func TestEncryptCompare(t *testing.T) {
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

	destination := make([]byte, len(plaintext))
	cipherBlock.Encrypt(destination, plaintext)

	t.Logf("encrypted plaintext: %x", destination)

	decrypted := make([]byte, len(destination))
	cipherBlock.Decrypt(decrypted, destination)

	for i, b := range decrypted {
		if b != plaintext[i] {
			t.Errorf("expeted plaintext and decrypted to equal, but got [%v] and [%v]", plaintext, decrypted)
		}
	}
}

func equal(a, b []uint32) bool {
	for i, x := range a {
		if b[i] != x {
			return false
		}
	}

	return true
}

// the following tests are from https://blog.nindalf.com/posts/implementing-aes/

func TestSubBypesNew(t *testing.T) {
	input := []uint32{0x8e9ff1c6, 0x4ddce1c7, 0xa158d1c8, 0xbc9dc1c9}
	expected := []uint32{0x19dba1b4, 0xe386f8c6, 0x326a3ee8, 0x655e78dd}

	if !equal(subBytes(input), expected) {
		t.Errorf("expected: %x - actual: %x", expected, subBytes(input))
	}
}

func TestShiftRowsNew(t *testing.T) {
	input := []uint32{
		0x8e9f01c6,
		0x4ddc01c6,
		0xa15801c6,
		0xbc9d01c6}
	expected := []uint32{
		0x8e9f01c6,
		0xdc01c64d,
		0x01c6a158,
		0xc6bc9d01}

	if !equal(shiftRows(input), expected) {
		t.Errorf("expected: %x - actual: %x", expected, shiftRows(input))
	}
}

func TestMixColumns(t *testing.T) {
	input := []uint32{
		0xdbf201c6,
		0x130a01c6,
		0x532201c6,
		0x455c01c6}
	expected := []uint32{
		0x8e9f01c6,
		0x4ddc01c6,
		0xa15801c6,
		0xbc9d01c6}

	if !equal(mixColumns(input), expected) {
		t.Errorf("expected: %x - actual: %x", expected, mixColumns(input))
	}
}
