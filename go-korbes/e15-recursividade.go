package main

// Recursividade

import "fmt"

// Minha solução

func main() {
	fmt.Println(fatorial(8))
}

var total int = 1

func fatorial(x int) int {

	if x > 0 {
		total = total * x
		fatorial(x - 1)
	}

	return total
}

// - Com recursividade. Go Playground: https://play.golang.org/p/ujsLnUhRp_
