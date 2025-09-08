package main

import "fmt"

func main() {
	fmt.Println(pgcd(10, 20))
	fmt.Println(pgcd(12, 24))
}

func pgcd(a, b int) int {
	if b == 0 {
		return a
	}
	return pgcd(b, a%b)
}
