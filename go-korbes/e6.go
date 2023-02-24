// - Utilizando como base o exercício anterior, utilize slicing para demonstrar os
// valores:
//     - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
//     - Do quinto ao último item do slice (incluindo o último item!)
//     - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
//     - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
//     - Desafio: obtenha o mesmo resultado acima utilizando a função len() para
// determinar o penúltimo item

package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)

	slice2 := slice[:3]
	fmt.Println(slice2)

	slice3 := slice[4:]
	fmt.Println(slice3)

	slice4 := slice[1:7]
	fmt.Println(slice4)

	slice5 := slice[2:(len(slice) - 1)]
	fmt.Println(slice5)
}