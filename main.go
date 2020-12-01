package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	for i := 1; i <= number; i++ {
		fmt.Println(fib(i))
	}

}
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
