package main

import "fmt"

func CountLetters(s string) int { //Argument en string et qui renvoie un interger
	var compteur int = 0          //On initie une variable compteur
	for j := 0; j < len(s); j++ { //Boucle for allant de 0 jusqu'a la longueur de s
		if s[j] >= 65 && s[j] <= 90 || s[j] >= 97 && s[j] <= 122 { //si j est entre 'a'  et 'z' ou j est entre 'A' et 'Z'
			compteur += 1 //on augmente le compteur de 1
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
