package main

// 10001번째 소수 구하기

func main() {
	var answer int = 0
	var cnt int = 1

	// Map 선언
	var prime map[int]bool
	prime = make(map[int]bool)
	prime[2] = true
	var isPrime bool = true

	// 숫자를 증가시키며
	// Map에 담은 소수들로 숫자를 나누어
	// 소수인지 확인
	for i := 3; ; i++ {
		for key := range prime {
			if i%key == 0 {
				isPrime = false
				break
			}
		}
		// 소수 일 경우 Map에 저장시키고
		// 10001번째면 프로그램 종료
		if isPrime {
			prime[i] = true
			cnt++
			if cnt == 10001 {
				answer = i
				break
			}
		}
		isPrime = true
	}

	println(answer)

}
