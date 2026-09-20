package main

import "fmt"

var idade int = 18

func main() {
	if idade >= 18 {
		fmt.Println("Maior de idade")
	} else {
		fmt.Println("Menor de idade")
	}
}
