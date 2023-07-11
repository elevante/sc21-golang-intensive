package app

import (
	"fmt"
	"math"
	"sort"
	"strings"
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
	if a.quantity%2 == 0 {
		medianValue = (a.Data[(a.quantity-1)/2] + a.Data[a.quantity/2]) / 2
	} else {
		medianValue = a.Data[a.quantity/2]
	}
	fmt.Printf("Median: %.2f\n", medianValue)
}

//ModeCalculation Мода – это величина признака (варианта), который наиболее часто встречается
//в данной совокупности,т.e. это варианта, имеющая наибольшую частоту.

func (a *Anscombe) ModeCalculation() {
	var arr []float64
	for i := 0; i < len(a.Data); i++ {
		for j := i + 1; j < len(a.Data); j++ {
			if a.Data[i] == a.Data[j] {
				arr = append(arr, a.Data[i])
			}
		}
	}
	if len(arr) == 0 {
		fmt.Println("No mode")
	} else {
		fmt.Printf("Mode: %v\n", strings.Trim(fmt.Sprintf("%.2f", arr), "[]"))
	}
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
	avarage := sum / float64(a.quantity)
	for _, n := range a.Data {
		standardDeviation += math.Pow(n-avarage, 2)
	}
	res := math.Sqrt(standardDeviation / float64(a.quantity))
	fmt.Printf("SD: %.2f\n", res)
}
