package eight

import (
	"fmt"
	"testing"
)

func Test_PartOne(t *testing.T) {
	if PartOne("input.txt") != fmt.Sprint(1965) {
		t.Fail()
	}
}

func Test_PartTwo(t *testing.T) {
	if PartTwo("input.txt") != fmt.Sprint("GZKJY") {
		t.Fail()
	}
}
