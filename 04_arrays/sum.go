package main

// Sum gets a slice of numbers and returns their sum.
func Sum(numbers []int) int {
	var r int
	for _, n := range numbers {
		r = r + n
	}
	return r
}

// SumAll gets a variable number of slices of numbers
// and returns the sum for each slice
func SumAll(slices ...[]int) []int {
	var r []int
	for _, s := range slices {
		r = append(r, Sum(s))
	}
	return r
}
