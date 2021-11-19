package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	sums := make([]int, 0)
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return sums
}

func SumAllTails(slices ...[]int) []int {
	sums := make([]int, 0)
	for _, slice := range slices {
		if len(slice) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := slice[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
