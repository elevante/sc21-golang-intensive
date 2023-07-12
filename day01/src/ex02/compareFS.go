package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type Data struct {
	arrOld []string
	arrNew []string
}

func Start() {
	flagNew := flag.String("new", "", "")
	flagOld := flag.String("old", "", "")

	flag.Parse()

	Compare(*flagOld, *flagNew)
}

func Compare(oldFile string, newFile string) {
	var d Data

	fileOld, err := os.Open(oldFile)
	if err != nil {
		log.Fatalf("Error when opening old file: %s", err)
	}

	fileOldScanner := bufio.NewScanner(fileOld)
	for i := 0; fileOldScanner.Scan(); i++ {
		d.arrOld = append(d.arrOld, fileOldScanner.Text())
	}
	defer func() {
		if err = fileOld.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fileNew, err := os.Open(newFile)
	if err != nil {
		log.Fatalf("Error when opening new file: %s", err)
	}

	fileNewScanner := bufio.NewScanner(fileNew)
	for i := 0; fileNewScanner.Scan(); i++ {
		d.arrNew = append(d.arrNew, fileNewScanner.Text())
	}
	defer func() {
		if err = fileNew.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for i := 0; i < len(d.arrNew); i++ {
		if Contains(d.arrOld, d.arrNew[i]) == false {
			fmt.Printf("ADDED %s\n", d.arrNew[i])
		}
	}

	for i := 0; i < len(d.arrOld); i++ {
		if Contains(d.arrNew, d.arrOld[i]) == false {
			fmt.Printf("REMOVED %s\n", d.arrOld[i])
		}
	}
}

func main() {
	Start()
}
