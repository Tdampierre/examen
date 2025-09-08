package main

import "fmt"

func main() {
	fmt.Println(pgdc(10, 20))
	fmt.Println(pgdc(12, 24))
}

func pgdc(a, b int) int {
	if b == 0 {
		return a
	}
	return pgdc(b, a%b)
}
