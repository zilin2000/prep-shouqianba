package sum

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (output []int) {
	lengthOfNumbers := len(numbersToSum)
	output = make([]int, lengthOfNumbers)

	for i, number := range numbersToSum {
		output[i] = Sum(number)
	}
	return
}
