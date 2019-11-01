package toolkit

import (
	"log"
	"testing"
)

func TestGenerateSectionIntSliceOfOrderly(t *testing.T) {
	log.Printf("Generate orderly slice:%+v\n", GenerateSectionIntSliceOfOrderly(1, 20, 3))
}

func TestGenerateSectionIntSliceOfDisorderly(t *testing.T) {
	log.Printf("Generate disorderly slice:%+v\n", GenerateSectionIntSliceOfDisorderly(15, 20))
}

func TestJoinItemOfStringSlice(t *testing.T) {
	log.Printf("joined string: %s\n", JoinItemOfStringSlice("", "中", "华", "人", "民", "共", "和", "国"))
}
