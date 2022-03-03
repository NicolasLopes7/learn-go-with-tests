package arrays

func Sum(arrayOfNumbers []int) (sum int) {
	for _, number := range arrayOfNumbers {
		sum += number
	}

	return
}
