package app

import (
	"fmt"
	"math"
	"sort"
)

//AverageCalculation Среднее арифметическое — разновидность среднего значения.
//Определяется как число,равное сумме всех чисел множества, делённой на их количество.

func (a *Anscombe) AverageCalculation() {
	var sum float64
	for _, n := range a.Data {
		sum += n
	}
	average := sum / float64(a.quantity)
	fmt.Printf("Mean: %.2f\n", average)
}

//MedianCalculation Медиана — число, которое находится в середине этого набора, если его упорядочить по возрастанию,
//то есть такое число, что половина из элементов набора не меньше него, а другая половина не больше.

func (a *Anscombe) MedianCalculation() {
	var medianValue float64
	sort.Slice(a.Data, func(i, j int) bool {
		return a.Data[i] < a.Data[j]
	})
	if len(a.Data) != 0 {
		if a.quantity%2 == 0 {
			medianValue = (a.Data[(a.quantity-1)/2] + a.Data[a.quantity/2]) / 2
		} else {
			medianValue = a.Data[a.quantity/2]
		}
	} else if len(a.Data) == 0 {
		fmt.Println("No data entered. Cannot calculate median.")
	}
	fmt.Printf("Median: %.2f\n", medianValue)
}

//ModeCalculation Мода – это величина признака (варианта), который наиболее часто встречается
//в данной совокупности,т.e. это варианта, имеющая наибольшую частоту.

func (a *Anscombe) ModeCalculation() {

	if len(a.Data) == 0 {
		fmt.Println("No mode")
		return
	}
	modeFrequency := make(map[float64]int)
	for _, val := range a.Data {
		modeFrequency[val]++
	}
	maxFrequency := 0
	var modes []float64

	for val, frequency := range modeFrequency {
		if frequency > maxFrequency {
			maxFrequency = frequency
			modes = []float64{val}
		} else if frequency == maxFrequency {
			modes = append(modes, val)
		}
	}
	if maxFrequency <= 1 {
		fmt.Println("No mode")
		return
	}

	minMode := modes[0]
	for _, mode := range modes {
		if mode < minMode {
			minMode = mode
		}
	}
	fmt.Printf("Mode: %.2f\n", minMode)
}

//StandardDeviationCalculation Среднеквадратическое отклонение — наиболее распространённый показатель рассеивания значений случайной величины
//относительно её математического ожидания (аналога среднего арифметического с бесконечным числом исходов).
//Обычно означает квадратный корень из дисперсии случайной величины, но иногда может означать тот или иной вариант оценки этого значения.

func (a *Anscombe) StandardDeviationCalculation() {
	var sum float64
	var standardDeviation float64
	for _, n := range a.Data {
		sum += n
	}
	average := sum / float64(a.quantity)
	for _, n := range a.Data {
		standardDeviation += math.Pow(n-average, 2)
	}
	res := math.Sqrt(standardDeviation / float64(a.quantity))
	fmt.Printf("SD: %.2f\n", res)
}
