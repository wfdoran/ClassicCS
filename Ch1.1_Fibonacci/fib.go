package main

import "fmt"

func fib2(n int) int {
	if n < 2 {
		return n
	}
	return fib2(n-1) + fib2(n-2)
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i, fib2(i))
	}
}
