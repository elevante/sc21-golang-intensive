package app

import (
	"bufio"
	"fmt"
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

		value, err := strconv.Atoi(input.Text())
		if err != nil {
			fmt.Println("Invalid value entered")
			break
		}

		if value >= -100000 && value <= 100000 {
			a.Data = append(a.Data, float64(value))
		}
	}
	a.quantity = len(a.Data)
}
