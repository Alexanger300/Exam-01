package main

import "fmt"

func CountLetters(s string) int {
	var compteur int = 0
	for j := 0; j < len(s); j++ {
		if s[j] >= 65 && s[j] <= 90 || s[j] >= 97 && s[j] <= 122 {
			compteur += 1
		}
	}
	return compteur
}
func main() {
	fmt.Println(CountLetters("Hello world!"))
	fmt.Println(CountLetters("123 456"))
	fmt.Println(CountLetters("GoLang1.21"))
	fmt.Println(CountLetters("école"))

}
