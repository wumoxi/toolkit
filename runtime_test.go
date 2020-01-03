package toolkit

import (
	"testing"
)

func TestMemAllocated(t *testing.T) {
	t.Logf("The current program footprint is: %s\n", MemAllocated())
}
