package main

// 1부터 20으로 나누어지는 공배수 찾기

func main() {

	var number int = 2521
	var isCorrect bool = true

	for {
		for i := 2; i <= 20; i++ {
			if number%i != 0 {
				isCorrect = false
				break
			}
		}
		if isCorrect {
			break
		}
		isCorrect = true
		number++
	}
	println(number)

}
