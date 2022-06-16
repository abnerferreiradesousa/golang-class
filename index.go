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

// Crie constantes tipadas e não-tipadas.

import "fmt"

const a int = 10
const b = 11


func main() {
	// Será que não estou redeclarando a variável b?
	b := "Olá amor"
	fmt.Printf("%T\n", a)
	fmt.Printf("%T", b)
}

// Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.

import "fmt"

const (
	_ = iota
	x = iota + 2022
	z
	y
	f
)

func main() {

	fmt.Println(x, z, y, f)
}

//Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
// Por exemplo:
  //  65
  //      U+0041 'A'
  //      U+0041 'A'
  //      U+0041 'A'
  //  66
  //      U+0042 'B'
  //      U+0042 'B'
  //      U+0042 'B' 
  //  ...e por aí vai.

	import "fmt"

	func main() {
		x := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	
		for i := 0; i < len(x); i++ {
			for index := 1; index <= 3; index++ {
				fmt.Printf("%#U\n", x[i])
			}
	
		}
	
	}