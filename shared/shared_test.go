package shared

import (
	"testing"
)

func TestMul(t *testing.T) {
	a := []byte{0x57, 0x57, 0x57, 0x57, 0x57}
	b := []byte{0x02, 0x04, 0x08, 0x10, 0x13}
	expected := []byte{0xae, 0x47, 0x8e, 0x07, 0xfe}
	for i := 0; i < len(a); i++ {
		out := FiniteFieldMultiply(a[i], b[i])
		if expected[i] != out {
			t.Errorf("mul(0x%02x, 0x%02x) = 0x%02x but got 0x%02x", a[i], b[i], expected[i], out)
		}
	}
}
