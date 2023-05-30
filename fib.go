package main

import (
	"fmt"
)

func fib(n int) <-chan int {
	result := make(chan int)

	go func() {
		result <- fib_internal(n-1) + fib_internal(n-2)
	}()

	return result
}

func fib_internal(n int) int {
	if n > 2 {
		return fib_internal(n-1) + fib_internal(n-2)
	} else if n < 1 {
		return 0
	} else {
		return 1
	}
}

func main() {
	i := 35
	s := 10

	for a := 0; a < s; a++ {
		fmt.Println(<-fib(i))
	}
}
