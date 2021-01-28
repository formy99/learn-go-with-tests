package main

func Sum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(nubersToSum ...[]int) []int {
	//lengthOfNumbers := len(nubersToSum)
	//sums = make([]int, lengthOfNumbers)
	//for i, numbers := range nubersToSum{
	//	sums[i] = Sum(numbers)
	//}
	var sums []int
	for _, numbers := range nubersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSumTail ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSumTail {
		if len(numbers) != 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}

	}
	return sums
}
