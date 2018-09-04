package main

func main() {
	println(isHappy(2))
}

func isHappy(n int) bool {
	m := make(map[int]bool)
	s := step(n)
	for s != 1 {
		if m[s] {
			return false
		}
		s = step(s)
		m[s] = true
	}
	return true
}

func step(n int) int {
	sum := 0
	for n > 0 {
		t := n%10
		sum += t * t
		n = n/10
	}
	return sum
}
