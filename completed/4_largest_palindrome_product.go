package main

import (
	"strconv"
)

func P4() int {
	a := 999
	b := 900
	var num int

	// var largest int

	for i := 0; i < 1000-b; i++ {
		b = 900
		for j := 0; j < 1000-b; j++ {
			num = a * b
			s := strconv.Itoa(num)
			if s[0] == s[5] && s[1] == s[4] && s[2] == s[3] {
				return num
			}
			b++
		}
		a--
	}
	return num
}
