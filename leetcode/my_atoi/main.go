package main

import "fmt"

func myAtoi(s string) int {
	maxInt := 2147483647
	minInt := -2147483648
	runes := []rune(s)
	n := len(runes)
	i := 0
	for i < n && runes[i] == ' ' {
		i++
	}

	sign := 1
	if i < n {
		if i < n && runes[i] == '-' {
			sign = -1
			i++
		} else if runes[i] == '+' {
			i++
		}
	}

	result := 0
	for i < n && runes[i] >= '0' && runes[i] <= '9' {
		digit := int(runes[i] - '0')
		if sign == 1 && (result > (maxInt-digit)/10) {
			return maxInt
		}
		if sign == -1 && (result > (-minInt-digit)/10) {
			return minInt
		}
		result = result*10 + digit
		i++
	}

	return sign * result
}

func main() {
	fmt.Println(myAtoi("   -042"))
}
