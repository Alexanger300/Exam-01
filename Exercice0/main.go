package main

import "fmt"

func PrintAlphabet() {
	for i := 'a'; i <= 'z'; i++ { //Boucle For allant de 'a' à 'z'
		fmt.Print(string(i))
	}
}
func main() { //Execution du programme
	PrintAlphabet()
}
