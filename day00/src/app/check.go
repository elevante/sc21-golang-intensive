package app

func CheckFlagsAndPrintCalculation() {
	flags := NewFlags()
	flags.Parse()

	resultAnscombe := Anscombe{}
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
	}
	if !*flags.Mean && !*flags.Median && !*flags.Mode && !*flags.SD {
		resultAnscombe.AverageCalculation()
		resultAnscombe.MedianCalculation()
		resultAnscombe.ModeCalculation()
		resultAnscombe.StandardDeviationCalculation()
	}
}
