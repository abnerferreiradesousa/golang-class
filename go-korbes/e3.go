// Crie constantes tipadas e não-tipadas.

package main

import "fmt"

const a int = 10
const b = 11

func main() {
	// Será que não estou redeclarando a variável b?
	b := "Olá amor"
	fmt.Printf("%T\n", a)
	fmt.Printf("%T", b)
}