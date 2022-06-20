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

func fib5(n int) int {
	if n == 0 {
		return 0
	}
	last := 0
	next := 1
	for i := 1; i < n; i++ {
		last, next = next, last+next
	}
	return next
}

func fib6(n int) <-chan int {
	ch := make(chan int)

	go func(ch chan int) {
		ch <- 0

		if n > 0 {
			last := 0
			next := 1
			for i := 1; i < n; i++ {
				last, next = next, last+next
				ch <- last
			}
		}
		close(ch)
	}(ch)

	return ch
}

func main() {
	memo = make(map[int]int)
	memo[0] = 0
	memo[1] = 1

	for i := 0; i < 20; i++ {
		fmt.Println(i, fib2(i), fib3(i), fib5(i))
	}

	for x := range fib6(20) {
		fmt.Println(x)
	}
}
