package array

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbersSlice := range numbersToSum {
		var i int
		for _, number := range numbersSlice {
			i += number
		}
		sums = append(sums, i)
		i = 0
	}
	return
}
