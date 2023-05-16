package main

import (
	"fmt"
	"godesde0/ejercicios"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()

	// estado, texto := variables.ConvierteaTexto(1234)

	// fmt.Println(estado)
	// fmt.Println(texto)

	/*if os := runtime.GOOS; os == "Linux." || os == "OS X." || os == "darwin" {
		fmt.Println("Esto no es Windows")
	} else {
		fmt.Println("Esto es Windows, es ", os)
	}

	switch os := runtime.GOOS; os {
	case "darwin1":
		fmt.Println("Esto es Darwin")
	case "Windows":
		fmt.Println("Esto es Windows")
	default:
		fmt.Printf("%s \n", os)
	}*/

	numero, mensaje := ejercicios.ConvierteaInt("19")
	fmt.Println(numero)
	fmt.Println(mensaje)

}
