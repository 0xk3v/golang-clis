package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	countBytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *countBytes))
}

func count(r io.Reader, countLines, countBytes bool) int {
	// Scanner
	scanner := bufio.NewScanner(r)

	if countLines {
		scanner.Split(bufio.ScanLines)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else {
		scanner.Split(bufio.ScanWords)
	}

	// Word count
	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
