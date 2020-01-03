package main

import (
	"strconv"
)

// Palindrome 찾기
// 3자리 수 x 3자리 수를 하여
// 가장 큰 Palindrome 찾기
// 999부터 역순으로 곱하여
// 해당 목적 값 찾기

func main() {
	var temp int = 0

	for i := 999; i >= 900; i-- {
		for j := 999; j >= 900; j-- {
			if isPalindrome(i * j) {
				if temp < i*j {
					temp = i * j
				}
			}
		}
	}

	println(temp)
}

func isPalindrome(sum int) bool {
	var str string
	var spl1 string
	var spl2 string

	str = strconv.Itoa(sum)
	spl1 = string(str[0]) + string(str[1]) + string(str[2])
	spl2 = string(str[5]) + string(str[4]) + string(str[3])

	if spl1 == spl2 {
		return true
	} else {
		return false
	}
}
