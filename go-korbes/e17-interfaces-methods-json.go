package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Nome string
	Age  int
}

type ByAge []Person

func (p ByAge) Len() int           { return len(p) }
func (p ByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByAge) Less(i, j int) bool { return p[i].Age < p[j].Age }

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
		{"Abner", 21},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}