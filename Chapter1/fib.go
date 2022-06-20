package main

import "fmt"

var memo map[int]int

func fib2(n int) int {
	if n < 2 {
		return n
	}
	return fib2(n-1) + fib2(n-2)
}

func fib3(n int) int {
	v, ok := memo[n]
	if ok {
		return v
	}
	memo[n] = fib3(n-1) + fib3(n-2)
	return memo[n]
}

func main() {
	memo = make(map[int]int)
	memo[0] = 0
	memo[1] = 1

	for i := 0; i < 20; i++ {
		fmt.Println(i, fib2(i), fib3(i))
	}
}
