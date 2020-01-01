package toolkit

import (
	"testing"
)

func TestGenerateSectionIntSliceOfOrderly(t *testing.T) {
	t.Logf("Generate orderly slice:%+v\n", GenerateSectionIntSliceOfOrderly(1, 20, 3))
}

func TestGenerateSectionIntSliceOfDisorderly(t *testing.T) {
	t.Logf("Generate disorderly slice:%+v\n", GenerateSectionIntSliceOfDisorderly(1, 20))
}

func TestJoinItemOfStringSlice(t *testing.T) {
	t.Logf("joined string: %s\n", JoinItemOfStringSlice("", "中", "华", "人", "民", "共", "和", "国"))
}

func TestGenerateSectionFibSlice(t *testing.T) {
	t.Logf("generate fib: %+v\n", GenerateSectionFibSlice(20))
}

func TestFactorial(t *testing.T) {
	t.Logf("%d is factorial value %+v\n", 10, Factorial(10))
}

func TestFactorialSliceOfUint64(t *testing.T) {
	t.Logf("generate factorial: %+v\n", FactorialSliceOfUint64(10))
}
