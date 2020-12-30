package seven

import (
	"fmt"
	"testing"
)

func Test_PartOne(t *testing.T) {
	if PartOne("input.txt") != fmt.Sprint(19650) {
		t.Fail()
	}
}

func Test_PartTwo(t *testing.T) {
	if PartTwo("input.txt") != fmt.Sprint(35961106) {
		t.Fail()
	}
}
