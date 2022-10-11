package main

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

*/

/*

SOLUTION: find GCD recursively
https://sites.math.rutgers.edu/~greenfie/gs2004/euclid.html

*/

func P5() int {
	found := false
	smallest := 2520
	for !found {
		if cascade(smallest, 20) {
			found = true
		} else {
			smallest += 10
		}
	}
	return smallest
}

// Cascades GCD until it equals 1
func cascade(x, y int) bool {
	if y == 1 {
		return true
	}
	z := gcd(x, y)
	if z != y {
		return false
	}
	return cascade(x, y-1)
}

// Calculates greatest common divisor
func gcd(a, b int) int {
	var d, r int
	if a < b {
		d = a
		r = b % a
	} else {
		d = b
		r = a % b
	}
	if r != 0 {
		d = gcd(d, r)
	}
	return d
}