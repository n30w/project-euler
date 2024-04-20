package main

func P6() int {
	var difference int
	// Sum of all rational numbers squared
	var squareOfSum int
	// Sum of every each squared number
	var sumOfSquares int

	for i := 1; i < 101; i++ {
		sumOfSquares += i * i
		squareOfSum += i
	}

	difference = (squareOfSum * squareOfSum) - sumOfSquares

	return difference
}

// Solution given by projecteuler.net at https://projecteuler.net/overview=006
func P6_solution() int {
	limit := 100
	sum := limit * (limit + 1) / 2
	sumSq := (((2 * limit) + 1) * (limit + 1) * limit) / 6
	return sum*sum - sumSq
}