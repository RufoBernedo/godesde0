package main

import (
	"fmt"
	"godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()

	estado, texto := variables.ConvierteaTexto(1234)

	fmt.Println(estado)
	fmt.Println(texto)
}
