package five

import "log"

type boardingPass struct {
	loc string
	row int
	col int
}

func (a boardingPass) CalcSeatID() int {
	return a.row*8 + a.col
}

type bySeatID []boardingPass

func (l bySeatID) Len() int           { return len(l) }
func (l bySeatID) Less(i, j int) bool { return l[i].CalcSeatID() < l[j].CalcSeatID() }
func (l bySeatID) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

type byLocation []boardingPass

func (l byLocation) Len() int { return len(l) }
func (l byLocation) Less(i, j int) bool {
	return l[i].row < l[j].row || (l[i].row == l[j].row && l[i].col < l[j].col)
}
func (l byLocation) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func parseBoardingPass(loc string) boardingPass {
	row := parseRowLocation(loc[:7])
	col := parseColumnLocation(loc[7:])
	return boardingPass{loc: loc, row: row, col: col}
}

func parseRowLocation(rowstr string) int {
	if len(rowstr) != 7 {
		log.Fatalf("Row location (%s) string must be 7 digits", rowstr)
	}
	row := 0
	poweroftwo := 1
	for i := 6; i >= 0; i-- {
		if rowstr[i] == 'B' {
			row += poweroftwo
		}
		poweroftwo *= 2
	}
	return row
}

func parseColumnLocation(colstr string) int {
	if len(colstr) != 3 {
		log.Fatalf("Column location (%s) string must be 3 digits", colstr)
	}
	col := 0
	poweroftwo := 1
	for i := 2; i >= 0; i-- {
		if colstr[i] == 'R' {
			col += poweroftwo
		}
		poweroftwo *= 2
	}
	return col
}
