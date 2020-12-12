package eleven

func initGrid(width, height int) [][]rune {
	grid := make([][]rune, width)
	for i := 0; i < width; i++ {
		grid[i] = make([]rune, height)
		for j := 0; j < height; j++ {
			grid[i][j] = 'x'
		}
	}
	return grid
}

func gridsAreEqual(grid1, grid2 [][]rune) bool {
	width := len(grid1)
	height := len(grid1[0])
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid1[x][y] != grid2[x][y] {
				return false
			}
		}
	}
	return true
}

func simulateSeats(grid [][]rune, visible bool) [][]rune {
	width := len(grid)
	height := len(grid[0])
	newgrid := initGrid(width, height)
	// For every cell in the grid
	for x, col := range grid {
		for y, seat := range col {
			// Skip if floor cell
			if seat == '.' {
				newgrid[x][y] = '.'
				continue
			}
			var occupied int
			if visible {
				occupied = countVisibleSeatsOccupied(grid, x, y)
			} else {
				occupied = countSurroundingSeatsOccupied(grid, x, y)
			}
			if seat == 'L' && occupied == 0 {
				newgrid[x][y] = '#'
				continue
			}
			maxOccupied := 4
			if visible {
				maxOccupied = 5
			}
			if seat == '#' && occupied >= maxOccupied {
				newgrid[x][y] = 'L'
				continue
			}
			newgrid[x][y] = grid[x][y]
		}
	}
	return newgrid
}

func countSurroundingSeatsOccupied(grid [][]rune, x int, y int) int {
	// Count all occupied seats surrounding position
	width := len(grid)
	height := len(grid[0])
	occupied := 0
	for i := x - 1; i <= x+1; i++ {
		if i < 0 || i >= width {
			continue
		}
		for j := y - 1; j <= y+1; j++ {
			if j < 0 || j >= height {
				continue
			}
			if !(i == x && j == y) && grid[i][j] == '#' {
				occupied++
			}
		}
	}
	return occupied
}

func countVisibleSeatsOccupied(grid [][]rune, x int, y int) int {
	// Count all occupied seats surrounding position
	width := len(grid)
	height := len(grid[0])
	occupied := 0
	for xOffset := -1; xOffset <= 1; xOffset++ {
		for yOffset := -1; yOffset <= 1; yOffset++ {
			i, j := x+xOffset, y+yOffset
			end := false
			if i < 0 || i >= width || j < 0 || j >= height {
				end = true
			}
			for !end && grid[i][j] == '.' {
				i += xOffset
				j += yOffset
				if i < 0 || i >= width || j < 0 || j >= height {
					end = true
				}
			}
			if end {
				continue
			}
			if !(i == x && j == y) && grid[i][j] == '#' {
				occupied++
			}
		}
	}
	return occupied
}

func countTotalOccupied(grid [][]rune) int {
	width := len(grid)
	height := len(grid[0])
	total := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid[x][y] == '#' {
				total++
			}
		}
	}
	return total
}
