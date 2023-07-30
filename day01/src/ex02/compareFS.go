package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func Start(oldFile string, newFile string) {
	var buffer []string

	buffer, err := readFileToBuffer(oldFile, buffer)
	if err != nil {
		log.Fatalf("%s", err)
	}

	compare(newFile, buffer, "REMOVED")

	buffer = nil

	buffer, err = readFileToBuffer(newFile, buffer)
	if err != nil {
		log.Fatalf("%s", err)
	}
	compare(oldFile, buffer, "ADDED")
}

func compare(filePath string, buffer []string, message string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error when opening new file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	for i := 0; fileScanner.Scan(); i++ {
		if Contains(buffer, fileScanner.Text()) == false {
			fmt.Printf("%s %s\n", message, fileScanner.Text())
		}
	}
}

func readFileToBuffer(filePath string, buffer []string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buffer = append(buffer, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return buffer, nil
}

func main() {
	flagNew := flag.String("new", "", "")
	flagOld := flag.String("old", "", "")
	flag.Parse()
	Start(*flagNew, *flagOld)
}
