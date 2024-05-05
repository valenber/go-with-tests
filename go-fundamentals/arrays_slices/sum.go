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

func SumAllTails(numberLists ...[]int) (sums []int) {
	for _, numbers := range numberLists {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
