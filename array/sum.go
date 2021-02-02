package array

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToTrail ...[]int) (sumtrail []int) {
	for _, numbers := range numbersToTrail {
		if len(numbers) != 0 {
			sumtrail = append(sumtrail, Sum(numbers[1:]))
		} else {
			sumtrail = append(sumtrail, 0)
		}
	}
	return
}
