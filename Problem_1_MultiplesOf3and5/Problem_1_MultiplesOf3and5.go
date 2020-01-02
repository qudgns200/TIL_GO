package main

// 1부터 1000 미만까지의 수에서
// 3 혹은 5로 나누어지는 수들의
// 합 구하기

func main() {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	println("sum = ", sum)
}
