package keyexpansion

import (
	"testing"

	"github.com/jackloughran/aes-go/shared"
)

func TestRotword(t *testing.T) {
	var word uint32 = 0x09cf4f3c
	var result uint32 = rotword(word)
	var expected uint32 = 0xcf4f3c09
	if expected != result {
		t.Errorf("Expected subword(%x) = %x but got %x", word, expected, result)
	}
}

func TestSubword(t *testing.T) {
	var word uint32 = 0xcf4f3c09
	var result uint32 = subword(word)
	var expected uint32 = 0x8a84eb01
	if expected != result {
		t.Errorf("Expected subword(%x) = %x but got %x", word, expected, result)
	}
}

func TestKeyExpansionAES128(t *testing.T) {
	Nk := 4
	Nr := 10
	key := []uint32{0x00010203, 0x04050607, 0x08090a0b, 0x0c0d0e0f}
	encKeys := keyExpansion(key, Nk, Nr)

	expectedEncKeys := [][]uint32{
		{0x00010203, 0x04050607, 0x08090a0b, 0x0c0d0e0f},
		{0xd6aa74fd, 0xd2af72fa, 0xdaa678f1, 0xd6ab76fe},
		{0xb692cf0b, 0x643dbdf1, 0xbe9bc500, 0x6830b3fe},
		{0xb6ff744e, 0xd2c2c9bf, 0x6c590cbf, 0x0469bf41},
		{0x47f7f7bc, 0x95353e03, 0xf96c32bc, 0xfd058dfd},
		{0x3caaa3e8, 0xa99f9deb, 0x50f3af57, 0xadf622aa},
		{0x5e390f7d, 0xf7a69296, 0xa7553dc1, 0x0aa31f6b},
		{0x14f9701a, 0xe35fe28c, 0x440adf4d, 0x4ea9c026},
		{0x47438735, 0xa41c65b9, 0xe016baf4, 0xaebf7ad2},
		{0x549932d1, 0xf0855768, 0x1093ed9c, 0xbe2c974e},
		{0x13111d7f, 0xe3944a17, 0xf307a78b, 0x4d2b30c5},
	}
	if !shared.Uint32MatrixEqual(expectedEncKeys, encKeys) {
		t.Error("Key expansion algorithm failed for encryption keys")
		t.Errorf("Expected \n%s\n but got\n %s", shared.Uint32MatrixToString(expectedEncKeys), shared.Uint32MatrixToString(encKeys))
	}
}
