// // Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.

package main

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