package arrays

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}

	return
}

func SumAll(numberLists ...[]int) (sums []int) {
	for _, numbers := range numberLists {
		sums = append(sums, Sum(numbers))
	}

	return
}
