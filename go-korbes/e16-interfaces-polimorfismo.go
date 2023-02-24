// - Crie um tipo "quadrado"
// - Crie um tipo "círculo"
// - Crie um método "área" para cada tipo que calcule e retorne a área da figura
//     - Área do círculo: 2 * π * raio
//     - Área do quadrado: lado * lado
// - Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
// - Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
// - Crie um valor de tipo "quadrado"
// - Crie um valor de tipo "círculo"
// - Use a função "info" para demonstrar a área do "quadrado"
// - Use a função "info" para demonstrar a área do "círculo"

package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	side float64
}

type circle struct {
	radius float64
}

func (p quadrado) area() float64 {
	return p.side * p.side
}

func (p circle) area() float64 {
	return 2 * math.Pi * p.radius
}

type figura interface {
	area() float64
}

func info(f figura) float64 {
	return f.area()
}

func main() {
	q := quadrado{17.4}
	c := circle{52.6}
	fmt.Println(math.Round(info(q)))
	fmt.Println(math.Round(info(c)))
}