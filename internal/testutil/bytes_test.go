package testutil

import (
	"fmt"
	"testing"
)

func TestRandomBytes(t *testing.T) {
	t.Parallel()

	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("%d random bytes", i), func(t *testing.T) {
			if got := RandomBytes(i); len(got) != i {
				t.Errorf("RandomBytes() got = %v, want %v", len(got), i)
				return
			}
		})
	}
}
