package nine

import (
	"fmt"
	"testing"
)

func Test_PartOne(t *testing.T) {
	if PartOne("input.txt", false) != fmt.Sprint(133015568) {
		t.Fail()
	}
}

func Test_PartTwo(t *testing.T) {
	if PartTwo("input.txt", false) != fmt.Sprint(16107959) {
		t.Fail()
	}
}
