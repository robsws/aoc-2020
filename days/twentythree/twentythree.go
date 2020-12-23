package twentythree

import (
	"fmt"
)

// PartOne - not yet implemented
func PartOne(filename string) string {
	circle := parseCups(filename, -1)
	for i := 0; i < 100; i++ {
		circle.doRound()
	}
	return circle.serialize()
}

// PartTwo - not yet implemented
func PartTwo(filename string) string {
	circle := parseCups(filename, 1000000)
	for i := 0; i < 10000000; i++ {
		circle.doRound()
	}
	onenode := circle.nodePointers[1]
	return fmt.Sprint(onenode.Next.Value * onenode.Next.Next.Value)
}
