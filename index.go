// Escreva um programa que mostre um número em decimal, binário e hexadecimal.

import "fmt"

func main() {
	a := 50
	fmt.Printf("%d, %#x, %b\n", a, a, a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%x", a)
}

// Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.

import "fmt"

func main() {
	a := 29
	b := 17

	z := a < b
	y := a > b
	u := a >= b
	v := a <= b
	x := a != b
	p := a == b

	fmt.Println(p, x, v, u, y, z)
}