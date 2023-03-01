package main

import "fmt"

func main() {
	// fatias(slices)
	arr := [4]string{"Hello", "Abner", "Sousa", "Ferreira"}
	fatiaoneway := arr[1:3]
	// append func
	appending := append(fatiaoneway, "H", "I")
	fmt.Printf("Resultado aqui: %v", appending)

	// mapas

	fmap := make(map[string]string)
	fmap["allfor"] = "one"

	fmt.Printf("Resultado aqui: %v", fmap["allfor"])

}
