package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func WordCount() {
	var files []string

	lFlag := flag.Bool("l", false, "Count line")
	wFlag := flag.Bool("w", false, "Count words")
	mFlag := flag.Bool("m", false, "Count characters")

	flag.Parse()

	files = flag.Args()

	for _, file := range files {

		fileFromSlice, err := os.Open(file)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}

		if *lFlag && strings.HasSuffix(file, ".txt") {
			scanner := bufio.NewScanner(fileFromSlice)

			countLine := 0
			for scanner.Scan() {
				countLine++
			}
			fmt.Printf("%d\t%s\n", countLine, file)
		}

		if *wFlag && strings.HasSuffix(file, ".txt") {

			scanner := bufio.NewScanner(fileFromSlice)
			wordsCount := 0

			for scanner.Scan() {
				line := scanner.Text()
				words := strings.Fields(line)
				wordsCount += len(words)
			}
			fmt.Printf("%d\t%s\n", wordsCount, file)

		}

		if *mFlag && strings.HasSuffix(file, ".txt") {
			
		}
	}

}

func main() {
	WordCount()
}
