package main

// 600851475143 의 가장 큰 소인수 찾기
// 주어진 숫자를 나누어 몫이 1이 될 때의
// 나눈 값을 리턴

func main() {
	var val int = 600851475143

	for i := 2; ; i++ {
		if val/i == 1 {
			break
		} else if val%i == 0 {
			val = val / i
		}
	}

	println(val)
}
