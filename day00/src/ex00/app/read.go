package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Anscombe struct {
	Data     []float64
	quantity int
}

func (a *Anscombe) Read() {

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {

		if input.Text() == "q" || input.Text() == "Q" {
			fmt.Println("Stop entering numbers")
			break
		}
		if input.Text() == "" {
			fmt.Println("Empty line. Nothing has been entered.")
		}

		value, err := strconv.Atoi(input.Text())
		if err != nil && input.Text() != "" {
			fmt.Println("Invalid value entered")
			continue
		}

		if value >= -100000 && value <= 100000 {
			a.Data = append(a.Data, float64(value))
		} else {
			log.Println("The entered value is greater than 100000 or less than -100000")
			os.Exit(1)
		}
	}
	a.quantity = len(a.Data)
}
