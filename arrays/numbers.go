package arrays

func Sum(numbers []int) (sum int) {
	for _, numbers := range numbers {
		sum += numbers
	}
	return sum
}
