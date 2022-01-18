package lib

func SumNumList(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}

	return total
}

//using recursion
func SumNumRecursion(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + SumNumRecursion(numbers[1:])
}
