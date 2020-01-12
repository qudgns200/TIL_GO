package main

// 500개 이상의 소인수를 갖는 숫자 찾기

func main() {

	var number int = 76576499 // 소인수 분해 할 수

	for {
		var quotient int = 0 // 소인수 분해 시 break를 위해 사용 될 변수
		var cnt int = 0      // 소인수 분해 개수
		for i := 1; i < number; i++ {
			if number%i == 0 {
				if quotient == i {
					break
				} else {
					quotient = number / i
					cnt++
				}
			}
		}
		if cnt >= 250 {
			break
		}
		number++
	}

	println(number)
}
