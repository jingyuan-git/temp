package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	res := 1.0
	x_contribute := x
	for n > 0 {
		if n%2 == 1 {
			res *= x_contribute
		}
		x_contribute *= x_contribute
		n /= 2
	}
	return res
}

func main() {
	fmt.Println(myPow(2.0, -2))
}
