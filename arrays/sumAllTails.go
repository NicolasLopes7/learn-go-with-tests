package arrays

func SumAllTails(numbersToSum ...[]int) (result []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			tail := numbers[1:]
			result = append(result, Sum(tail))
		}
	}

	return
}
