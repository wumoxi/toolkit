package toolkit

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	fmt.Printf("Generate random integer number: %d\n", Random(100000, 999999))
}
