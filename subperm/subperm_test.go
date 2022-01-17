package subperm

import (
	"testing"
)

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

	if !equal(SubBytes(input), expected) {
		t.Errorf("expected: %x - actual: %x", expected, SubBytes(input))
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

	if !equal(ShiftRows(input), expected) {
		t.Errorf("expected: %x - actual: %x", expected, ShiftRows(input))
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

	if !equal(MixColumns(input), expected) {
		t.Errorf("expected: %x - actual: %x", expected, MixColumns(input))
	}
}
