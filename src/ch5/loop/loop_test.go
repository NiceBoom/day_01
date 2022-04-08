package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(i)
	}
}
