package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var files []string
	var sl []string

	FlagA := flag.Bool("a", false, "")

	flag.Parse()

	files = flag.Args()

	if !*FlagA {
		fmt.Println(files)

		for _, file := range files {
			first := strings.Split(file, ".")
			fmt.Println(first[0])
			sec := time.Now().Unix()
			newFile := first[0] + "_" + strconv.FormatInt(sec, 10) + "." + "tag.gz"
			sl = append(sl, newFile)
			fmt.Println(sl)
		}

		for _, file := range sl {
			fileO, err := os.Create(file)
			if err != nil { // если возникла ошибка
				fmt.Println("Unable to create file:", err)
				os.Exit(1) // выходим из программы
			}
			defer fileO.Close() // закрываем файл
		}
	}

}
