// - Usando uma literal composta:
//      - Crie um array que suporte 5 valores to tipo int
//      - Atribua valores aos seus índices
// - Utilize range e demonstre os valores do array.
// - Utilizando format printing, demonstre o tipo do array.

package main

import "fmt"

func main() {

	slice := [5]int{1, 2, 3, 4, 5}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", slice)

}