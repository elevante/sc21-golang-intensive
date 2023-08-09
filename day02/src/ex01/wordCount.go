package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

func countLine(sc *bufio.Scanner, file string, wg *sync.WaitGroup) {
	defer wg.Done()

	countLine := 0
	for sc.Scan() {
		countLine++
	}
	fmt.Printf("%d\t%s\n", countLine, file)
}

func countWord(sc *bufio.Scanner, file string, wg *sync.WaitGroup) {
	defer wg.Done()

	wordsCount := 0
	for sc.Scan() {
		line := sc.Text()
		words := strings.Fields(line)
		wordsCount += len(words)
	}
	fmt.Printf("%d\t%s\n", wordsCount, file)
}

func countCharacter(sc *bufio.Scanner, file string, wg *sync.WaitGroup) {
	defer wg.Done()

	charCount := 0
	for sc.Scan() {
		line := sc.Text()
		lineWithoutSpaces := strings.ReplaceAll(line, " ", "")
		charCount += utf8.RuneCountInString(lineWithoutSpaces)
	}
	fmt.Printf("%d\t%s\n", charCount, file)
}

func worker() {
	var files []string

	lFlag := flag.Bool("l", false, "Count line")
	wFlag := flag.Bool("w", false, "Count words")
	mFlag := flag.Bool("m", false, "Count characters")

	flag.Parse()

	files = flag.Args()

	var wg sync.WaitGroup

	for _, file := range files {

		fileFromSlice, err := os.Open(file)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}

		defer fileFromSlice.Close()

		scanner := bufio.NewScanner(fileFromSlice)

		wg.Add(1)
		if *lFlag && strings.HasSuffix(file, ".txt") {
			go countLine(scanner, file, &wg)
		}

		if *wFlag && strings.HasSuffix(file, ".txt") {
			go countWord(scanner, file, &wg)
		}

		if *mFlag && strings.HasSuffix(file, ".txt") {
			go countCharacter(scanner, file, &wg)
		}

		if !*mFlag && !*lFlag && !*wFlag {
			go countWord(scanner, file, &wg)
		}
	}
	wg.Wait()
}

func main() {
	worker()
}
