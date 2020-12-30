package five

import (
	"fmt"
	"testing"
)

func Test_PartOne(t *testing.T) {
	if PartOne("input.txt") != fmt.Sprint(15426686) {
		t.Fail()
	}
}

func Test_PartTwo(t *testing.T) {
	if PartTwo("input.txt") != fmt.Sprint(11430197) {
		t.Fail()
	}
}
