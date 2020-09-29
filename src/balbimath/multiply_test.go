package balbimath

import "testing"

func TestMultiply(t *testing.T) {
	if (Multiply(1, 1) != 1) {
		t.Error("Multiplicar 1 * 1 deve ser 1")
	}
}