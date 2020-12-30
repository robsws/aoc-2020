package six

import (
	"fmt"
	"testing"
)

func Test_PartOne(t *testing.T) {
	if PartOne("input.txt") != fmt.Sprint(249308) {
		t.Fail()
	}
}

func Test_PartTwo(t *testing.T) {
	if PartTwo("input.txt") != fmt.Sprint(349) {
		t.Fail()
	}
}
