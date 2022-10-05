package main

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

const BIGASSNUMBER = 600_851_475_143

func P3() int {
	i := 2
	num := BIGASSNUMBER

	for num > 1 {
		if num%i == 0 {
			num /= i
		} else {
			i++
		}
	}

	return i
}
