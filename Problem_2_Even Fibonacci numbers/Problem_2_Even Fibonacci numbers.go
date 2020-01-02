package main

// 피보나치 구하기
// 수열이 4,000,000을 넘으면
// 짝수의 합 리턴

func main() {
	var num1 int = 1
	var num2 int = 2
	var sum int = num2
	var temp int = num2
	for i := 0; ; i++ {
		temp = num2
		num2 = num1 + num2

		if num2 > 4000000 {
			break
		} else {
			if num2%2 == 0 {
				sum += num2
			}
		}
		num1 = temp
	}
	println("sum = ", sum)
}
