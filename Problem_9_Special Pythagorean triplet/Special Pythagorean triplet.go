package main

// 피타고라스 수 구하기
// a + b + c = 1000
// a² + b² = c²

func main() {
	var sum int = 1000
	var a, b, c int
	var isTrue bool = false
	for b = 1; b < sum; b++ {
		for a = 1; a < b; a++ {
			c = sum - a - b
			if (b < c) && (c*c == a*a+b*b) {
				isTrue = true
				break
			}
		}
		if isTrue {
			break
		}
	}
	println(a * b * c)
}
