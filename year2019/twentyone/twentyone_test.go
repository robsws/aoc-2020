package twentyone

import (
	"fmt"
	"testing"
)

func Test_PartOne(t *testing.T) {
	if PartOne("input.txt") != fmt.Sprint(2412) {
		t.Fail()
	}
}

func Test_PartTwo(t *testing.T) {
	if PartTwo("input.txt") != "mfp,mgvfmvp,nhdjth,hcdchl,dvkbjh,dcvrf,bcjz,mhnrqp" {
		t.Fail()
	}
}
