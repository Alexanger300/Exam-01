package main

import "fmt"

func PGCD(nbr int, nbr2 int) int {
	var test int = 0
	for i := 1; i <= 9; i++ {
		if nbr%i == 0 && nbr2%i == 0 {
			test = i
		}
	}
	return test
}

func main() {
	fmt.Println(PGCD(6, 12))
	fmt.Println(PGCD(20, 48))
}
