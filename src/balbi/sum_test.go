package balbi;

import "testing"

func TestSum(t *testing.T) {
	if (Sum(1, 2) != 5) {
		t.Error("Soma errada")
	}
}