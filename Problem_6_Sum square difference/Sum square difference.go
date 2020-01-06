package main

// 1부터 100까지 각각의 제곱수의 합과
// 1부터 100까지 합의 제곱수의 차이 구하기

func main() {

	var sum1 int = 0
	var sum2 int = 0

	for i := 1; i <= 100; i++ {
		sum1 += i * i
		sum2 += i
	}

	sum2 = sum2 * sum2

	println(sum2 - sum1)

}
