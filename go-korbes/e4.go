// //Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
// // Por exemplo:
//   //  65
//   //      U+0041 'A'
//   //      U+0041 'A'
//   //      U+0041 'A'
//   //  66
//   //      U+0042 'B'
//   //      U+0042 'B'
//   //      U+0042 'B'
//   //  ...e por aí vai.

package main

import "fmt"

	func main() {
		x := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

		for i := 0; i < len(x); i++ {
			for index := 1; index <= 3; index++ {
				fmt.Printf("%#U\n", x[i])
			}

		}

	}