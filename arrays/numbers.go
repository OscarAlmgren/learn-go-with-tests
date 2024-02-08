package arrays

func Sum(numbers []int) (sum int) {
	for _, numbers := range numbers {
		sum += numbers
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
