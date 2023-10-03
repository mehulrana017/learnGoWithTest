package main

func SumAll(numberToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
