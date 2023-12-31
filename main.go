package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Define command-line flags
	var printBytes bool
	var printLines bool
	var printWords bool

	flag.BoolVar(&printBytes, "c", false, "print bytes")
	flag.BoolVar(&printLines, "l", false, "print lines")
	flag.BoolVar(&printWords, "w", false, "print words")

	flag.Parse()

	var data []byte
	var fileName string

	// Check if a filename is provided as a command-line argument
	if flag.NArg() > 0 {
		fileName = flag.Arg(0)

		// Reads the content of the file specified
		fileData, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal("Error: ", err)
			os.Exit(1)
		}
		data = fileData
	} else {
		// Read data from stdin
		stdinData, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal("Error reading from stdin: ", err)
			os.Exit(1)
		}
		data = stdinData
	}

	// If none of the flags is set, print all counts
	if !printBytes && !printLines && !printWords {
		printCounts(data, fileName)
	}

	// Print bytes if the flag is set
	if printBytes {
		printResult(countBytes(data))
	}

	// Print lines if the flag is set
	if printLines {
		printResult(countLines(data))
	}

	// Print words if the flag is set
	if printWords {
		printResult(countWords(data))
	}
}

func printCounts(data []byte, fileName string) {
	numBytes := countBytes(data)
	numLines := countLines(data)
	numWords := countWords(data)

	fmt.Printf("%d %d %d %s\n", numBytes, numLines, numWords, fileName)
}

func printResult(count int) {
	fmt.Println(count)
}

func countBytes(data []byte) int {
	return len(data)
}

func countLines(data []byte) int {
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}

func countWords(data []byte) int {
	content := string(data)
	words := strings.Fields(content)
	return len(words)
}
