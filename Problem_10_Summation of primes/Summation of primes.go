package main

// 2,000,000 이하의 소수들의 합 구하기

func main() {
	var answer int = 2

	// Map 선언
	var prime map[int]bool
	prime = make(map[int]bool)
	prime[2] = true
	var isPrime bool = true

	// 숫자를 증가시키며
	// Map에 담은 소수들로 숫자를 나누어
	// 소수인지 확인
	for i := 3; i <= 2000000; i++ {
		for key := range prime {
			if i%key == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			prime[i] = true
			answer += i
		}
		isPrime = true
		println(i)
	}

	println(answer)
}
