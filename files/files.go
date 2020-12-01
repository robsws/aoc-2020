package files

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// StreamLines - Read lines of file and write them to out channel
func StreamLines(filename string, out chan string) {
	buf, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		close(out)
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	reader := bufio.NewScanner(buf)
	for reader.Scan() {
		out <- reader.Text()
	}

	err = reader.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// StreamInts - Stream a file of integers
func StreamInts(filename string, out chan int) {
	fileStream := make(chan string)
	defer close(out)
	go StreamLines(filename, fileStream)
	for line := range fileStream {
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		out <- i
	}
}
