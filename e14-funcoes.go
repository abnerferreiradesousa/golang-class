package main

// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.

import "fmt"

func main() {
	slice := []int{4, 3, 4, 4, 5, 6, 7, 8, 9, 10}
	i := fatorial(slice...)
	t := fatorialcomslice(slice)
	fmt.Println(i)
	fmt.Println(t)
}

func fatorial(x ...int) int {
	total := 0
	for _, v := range x {
		total += v

	}
	return total
}

func fatorialcomslice(x []int) int {
	total := 0
	for _, v := range x {
		total += v

	}
	return total
}