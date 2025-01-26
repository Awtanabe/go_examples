package main


func Sum(numbers []int) int {

	sum := 0 
	for _ ,num := range numbers {
		sum += num 
	} 
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lens := len(numbersToSum)
	sums := make([]int, lens)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}