package twenty

import (
	"aoc-2020/files"
	"aoc-2020/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

// PartOne - find the 4 corner tiles and multiply their IDs
func PartOne(filename string) string {
	tiles := parseTiles(filename)
	corners := findCornerTiles(tiles)
	result := 1
	for _, t := range corners {
		result *= t.ID
	}
	return fmt.Sprint(result)
}

// PartTwo - not yet implemented
func PartTwo(filename string) string {
	tiles := parseTiles(filename)
	image := assembleTiles(tiles)
	amount := findSeaMonsters(image)
	pixelsOn := countPixelsOn(image)
	// 15 pixels in sea monster
	return fmt.Sprint(pixelsOn - amount*15)
}

func parseTiles(filename string) []tile {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	headerRe := regexp.MustCompile("^Tile ([0-9]+):$")
	tiles := make([]tile, 0)
	// Parse the input
	for line := range fileStream {
		if line == "" {
			continue
		}
		submatches := headerRe.FindStringSubmatch(line)
		tileID := utils.MustAtoi(submatches[1])
		t := makeTile(tileID)
		for i := 0; i < 10; i++ {
			gridLine := <-fileStream
			// Convert the pixels input to bool 2D array
			pixelsSlice := strings.Split(gridLine, "")
			pixels := [10]bool{}
			for j, pixel := range pixelsSlice {
				pixels[j] = pixel == "#"
			}
			t.Grid[i] = pixels
		}
		tiles = append(tiles, t)
	}
	return tiles
}

func assembleTiles(tiles []tile) [][]bool {
	// Init tile arrangement slice
	width := int(math.Sqrt(float64(len(tiles))))
	arrangement := make([][]tile, width)
	for i := 0; i < width; i++ {
		arrangement[i] = make([]tile, width)
	}
	// Find and place top left corner to start
	corners := findCornerTiles(tiles)
	arrangement[0][0] = corners[0]
	/* Find and orient tiles, starting by fitting the first tile
	   in each row underneath the one above, then fitting the rest
	   of the row to the right, for every row. */
	fittedTiles := utils.MakeSet()
	fittedTiles.Add(arrangement[0][0].ID)
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			if y == 0 && x == 0 {
				// We've already fit the top corner into the grid.
				continue
			} else if x == 0 {
				// Fit a tile underneath the one above for the first in the row.
				for _, t := range tiles {
					if fittedTiles.Contains(t.ID) {
						continue
					}
					fits := arrangement[y-1][0].fitUnderneath(&t)
					if fits {
						arrangement[y][0] = t
						fittedTiles.Add(t.ID)
						break
					}
				}
				continue
			}
			// For all other spaces in the row, fit a tile to the right
			// of the previous one.
			for _, t := range tiles {
				if fittedTiles.Contains(t.ID) {
					continue
				}
				fits := arrangement[y][x-1].fitToRight(&t)
				if fits {
					arrangement[y][x] = t
					fittedTiles.Add(t.ID)
					break
				}
			}
		}
	}
	/* Tiles are assembled in correct order and orientation. Now
	   trim away the borders and format as a single image. */
	image := make([][]bool, width*8)
	for i := 0; i < width*8; i++ {
		image[i] = make([]bool, width*8)
	}
	for ay := 0; ay < len(arrangement); ay++ {
		for ax := 0; ax < len(arrangement[0]); ax++ {
			subimage := arrangement[ay][ax].trimBorders()
			for iy := 0; iy < 8; iy++ {
				for ix := 0; ix < 8; ix++ {
					image[ay*8+iy][ax*8+ix] = subimage[iy][ix]
				}
			}
		}
	}
	return image
}

func findSeaMonsters(image [][]bool) int {
	x := true
	o := false
	seaMonster := [3][20]bool{
		{o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, x, o},
		{x, o, o, o, o, x, x, o, o, o, o, x, x, o, o, o, o, x, x, x},
		{o, x, o, o, x, o, o, x, o, o, x, o, o, x, o, o, x, o, o, o},
	}
	orientations := [8]func() [][]bool{
		func() [][]bool { return image },
		func() [][]bool { return flipImageVert(image) },
		func() [][]bool { return rotateImage(image) },
		func() [][]bool { return flipImageVert(rotateImage(image)) },
		func() [][]bool { return rotateImage(rotateImage(image)) },
		func() [][]bool { return flipImageVert(rotateImage(rotateImage(image))) },
		func() [][]bool { return rotateImage(rotateImage(rotateImage(image))) },
		func() [][]bool { return flipImageVert(rotateImage(rotateImage(rotateImage(image)))) },
	}
	amount := 0
	for _, orientationFunc := range orientations {
		orientation := orientationFunc()
		for y := 0; y < len(orientation)-3; y++ {
			for x := 0; x < len(orientation)-20; x++ {
				found := true
				for sy := 0; sy < 3; sy++ {
					for sx := 0; sx < 20; sx++ {
						if seaMonster[sy][sx] && !orientation[y+sy][x+sx] {
							found = false
							break
						}
					}
					if !found {
						break
					}
				}
				if found {
					amount++
				}
			}
		}
		if amount > 0 {
			break
		}
	}
	return amount
}

func (t tile) fitToRight(right *tile) bool {
	if t.ID == right.ID {
		return false
	}
	lb := t.getBorders()
	rb := right.getBorders()
	if lb.Right == rb.Left {
		return true
	}
	if lb.Right == flipRow(rb.Left) {
		right.flipVert()
		return true
	}
	if lb.Right == rb.Bottom {
		right.rotate(1)
		return true
	}
	if lb.Right == flipRow(rb.Bottom) {
		right.flipHoriz()
		right.rotate(1)
		return true
	}
	if lb.Right == rb.Right {
		right.flipHoriz()
		return true
	}
	if lb.Right == flipRow(rb.Right) {
		right.flipVert()
		right.flipHoriz()
		return true
	}
	if lb.Right == rb.Top {
		right.rotate(1)
		right.flipHoriz()
		return true
	}
	if lb.Right == flipRow(rb.Top) {
		right.flipHoriz()
		right.rotate(1)
		right.flipHoriz()
		return true
	}
	return false
}

func (t tile) fitUnderneath(under *tile) bool {
	if t.ID == under.ID {
		return false
	}
	tb := t.getBorders()
	ub := under.getBorders()
	if tb.Bottom == ub.Top {
		return true
	}
	if tb.Bottom == flipRow(ub.Top) {
		under.flipHoriz()
		return true
	}
	if tb.Bottom == ub.Left {
		under.rotate(1)
		under.flipHoriz()
		return true
	}
	if tb.Bottom == flipRow(ub.Left) {
		under.flipVert()
		under.rotate(1)
		under.flipHoriz()
		return true
	}
	if tb.Bottom == ub.Bottom {
		under.flipVert()
		return true
	}
	if tb.Bottom == flipRow(ub.Bottom) {
		under.flipHoriz()
		under.flipVert()
		return true
	}
	if tb.Bottom == ub.Right {
		under.rotate(1)
		under.flipHoriz()
		under.flipVert()
		return true
	}
	if tb.Bottom == flipRow(ub.Right) {
		under.rotate(1)
		under.flipVert()
		return true
	}
	return false
}

type tile struct {
	ID   int
	Grid [10][10]bool
}

type borders struct {
	Top    [10]bool
	Bottom [10]bool
	Left   [10]bool
	Right  [10]bool
}

func makeTile(ID int) tile {
	return tile{ID, [10][10]bool{}}
}

func (t tile) getBorders() borders {
	b := borders{}
	b.Top = t.Grid[0]
	b.Bottom = t.Grid[9]
	leftBorder := [10]bool{}
	rightBorder := [10]bool{}
	for i := 0; i < 10; i++ {
		leftBorder[i] = t.Grid[i][0]
		rightBorder[i] = t.Grid[i][9]
	}
	b.Left = leftBorder
	b.Right = rightBorder
	return b
}

func (t *tile) rotate(times int) {
	for i := 0; i < times; i++ {
		newGrid := [10][10]bool{}
		newX := 9
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				newGrid[x][newX] = t.Grid[y][x]
			}
			newX--
		}
		t.Grid = newGrid
	}
}

func (t *tile) flipHoriz() {
	newGrid := [10][10]bool{}
	for y := 0; y < 10; y++ {
		newGrid[y] = flipRow(t.Grid[y])
	}
	t.Grid = newGrid
}

func (t *tile) flipVert() {
	newGrid := [10][10]bool{}
	for y := 0; y < 10; y++ {
		newGrid[9-y] = t.Grid[y]
	}
	t.Grid = newGrid
}

func (t tile) trimBorders() [8][8]bool {
	image := [8][8]bool{}
	for y := 1; y < 9; y++ {
		for x := 1; x < 9; x++ {
			image[y-1][x-1] = t.Grid[y][x]
		}
	}
	return image
}

func (t tile) print() {
	fmt.Println()
	for y := 0; y < len(t.Grid); y++ {
		for x := 0; x < len(t.Grid[0]); x++ {
			if t.Grid[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func flipRow(b [10]bool) [10]bool {
	result := [10]bool{}
	for i := 0; i < 10; i++ {
		result[9-i] = b[i]
	}
	return result
}

func findCornerTiles(tiles []tile) []tile {
	borderCounts := countBorders(tiles)
	imageEdgeBorders := utils.MakeSet()
	for border, n := range borderCounts {
		if n == 1 {
			imageEdgeBorders.Add(border)
		}
	}
	// find the 4 corner tiles with 4 unique borders and orient them the same
	corners := make([]tile, 0, 4)
	for _, t := range tiles {
		b := t.getBorders()
		if imageEdgeBorders.Contains(b.Top) && imageEdgeBorders.Contains(b.Left) {
			corners = append(corners, t)
			continue
		}
		if imageEdgeBorders.Contains(b.Bottom) && imageEdgeBorders.Contains(b.Left) {
			t.flipVert()
			corners = append(corners, t)
			continue
		}
		if imageEdgeBorders.Contains(b.Bottom) && imageEdgeBorders.Contains(b.Right) {
			t.flipVert()
			t.flipHoriz()
			corners = append(corners, t)
			continue
		}
		if imageEdgeBorders.Contains(b.Top) && imageEdgeBorders.Contains(b.Right) {
			t.flipHoriz()
			corners = append(corners, t)
			continue
		}
	}
	return corners
}

func countBorders(tiles []tile) map[[10]bool]int {
	borderCounts := make(map[[10]bool]int)
	for _, t := range tiles {
		borders := t.getBorders()
		borderCounts[borders.Top]++
		borderCounts[flipRow(borders.Top)]++
		borderCounts[borders.Bottom]++
		borderCounts[flipRow(borders.Bottom)]++
		borderCounts[borders.Left]++
		borderCounts[flipRow(borders.Left)]++
		borderCounts[borders.Right]++
		borderCounts[flipRow(borders.Right)]++
	}
	return borderCounts
}

func rotateImage(image [][]bool) [][]bool {
	newImage := make([][]bool, len(image))
	for i := 0; i < len(image); i++ {
		newImage[i] = make([]bool, len(image))
	}
	newX := len(image) - 1
	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image); x++ {
			newImage[x][newX] = image[y][x]
		}
		newX--
	}
	return newImage
}

func flipImageVert(image [][]bool) [][]bool {
	newImage := make([][]bool, len(image))
	for i, row := range image {
		newImage[len(image)-(i+1)] = row
	}
	return newImage
}

func countPixelsOn(image [][]bool) int {
	amount := 0
	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image[0]); x++ {
			if image[y][x] {
				amount++
			}
		}
	}
	return amount
}

func printImage(image [][]bool) {
	fmt.Println()
	for y := 0; y < len(image); y++ {
		if y%8 == 0 {
			fmt.Println()
		}
		for x := 0; x < len(image[0]); x++ {
			if x%8 == 0 {
				fmt.Print(" ")
			}
			if image[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()
}
