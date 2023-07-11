package flag

import (
	"golang-indermediate-intensive/module00/ex00/app/anscombe"
)

func CheckFlagsAndPrintCalculation() {
	flags := NewFlags()
	flags.Parse()

	resultAnscombe := anscombe.Anscombe{}
	resultAnscombe.Read()

	if *flags.Mean {
		resultAnscombe.AverageCalculation()
	}
	if *flags.Median {
		resultAnscombe.MedianCalculation()
	}
	if *flags.Mode {
		resultAnscombe.ModeCalculation()
	}
	if *flags.SD {
		resultAnscombe.StandardDeviationCalculation()
	} else {
		resultAnscombe.AverageCalculation()
		resultAnscombe.MedianCalculation()
		resultAnscombe.ModeCalculation()
		resultAnscombe.StandardDeviationCalculation()
	}
}
