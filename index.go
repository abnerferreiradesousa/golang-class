// Escreva um programa que mostre um número em decimal, binário e hexadecimal.

import "fmt"

func main() {
	a := 50
	fmt.Printf("%d, %#x, %b\n", a, a, a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%x", a)
}